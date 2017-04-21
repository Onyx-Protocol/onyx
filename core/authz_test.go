package core

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"chain/core/accesstoken"
	"chain/database/pg/pgtest"
	"chain/database/raft"
)

func TestAuthz(t *testing.T) {
	ctx := context.Background()
	_, db := pgtest.NewDB(t, pgtest.SchemaPath)
	accessTokens := &accesstoken.CredentialStore{db}

	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	raftDir := filepath.Join(currentDir, "/.testraft")
	err = os.Mkdir(raftDir, 0700)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(raftDir)

	raftDB, err := raft.Start("", raftDir, "", false)
	if err != nil {
		t.Fatal(err)
	}
	api := &API{
		mux:          http.NewServeMux(),
		raftDB:       raftDB,
		accessTokens: accessTokens,
	}
	api.buildHandler()
	server := httptest.NewServer(api)
	defer server.Close()

	testPolicies := []string{
		"client-readwrite",
		"client-readonly",
		"network",
		"monitoring",
		"internal",
		"public",
	}
	tokens := make(map[string]*accesstoken.Token)
	for i := 0; i < len(testPolicies); i++ {
		token, err := accessTokens.Create(ctx, fmt.Sprintf("token%d", i), "")
		if err != nil {
			t.Fatal(err)
		}
		tokens[policies[i]] = token
		err = api.createGrant(ctx, apiGrant{
			GuardType: "access_token",
			GuardData: map[string]interface{}{"id": token.ID},
			Policy:    policies[i],
		})
		if err != nil {
			t.Fatal(err)
		}
	}

	pathTokenMap := map[string]map[string]bool{
		"/create-account": map[string]bool{
			"client-readwrite": true,
			"client-readonly":  false,
			"network":          false,
			"monitoring":       false,
			"internal":         false,
			"public":           false,
		},
		"/list-accounts": map[string]bool{
			"client-readwrite": true,
			"client-readonly":  true,
			"network":          false,
			"monitoring":       false,
			"internal":         false,
			"public":           false,
		},
		"/reset": map[string]bool{
			"client-readwrite": true,
			"client-readonly":  false,
			"network":          false,
			"monitoring":       false,
			"internal":         true,
			"public":           false,
		},
		networkRPCPrefix + "get-block": map[string]bool{
			"client-readwrite": false,
			"client-readonly":  false,
			"network":          true,
			"monitoring":       false,
			"internal":         false,
			"public":           false,
		},
		"/info": map[string]bool{
			"client-readwrite": true,
			"client-readonly":  true,
			"network":          true,
			"monitoring":       true,
			"internal":         false,
			"public":           false,
		},
		"/debug/pprof/symbol": map[string]bool{
			"client-readwrite": true,
			"client-readonly":  true,
			"network":          false,
			"monitoring":       true,
			"internal":         false,
			"public":           false,
		},
		// "/raft/msg": map[string]bool{
		// 	"client-readwrite": false,
		// 	"client-readonly":  false,
		// 	"network":          false,
		// 	"monitoring":       false,
		// 	"internal":         true,
		// 	"public":           false,
		// },
		"/docs": map[string]bool{ // public is open to all
			"client-readwrite": true,
			"client-readonly":  true,
			"network":          true,
			"monitoring":       true,
			"internal":         true,
			"public":           true,
		},
	}

	for path, policyMap := range pathTokenMap {
		for policy, want := range policyMap {
			got := tryEndpoint(t, server.URL, path, tokens[policy])
			if got != want {
				t.Errorf("auth(%s, %s) = %t want %t", path, policy, got, want)
			}
		}
	}
}

func tryEndpoint(t testing.TB, baseURL, path string, token *accesstoken.Token) bool {
	req, err := http.NewRequest("POST", baseURL+path, bytes.NewReader([]byte("{}")))
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	req.SetBasicAuth(token.ID, strings.Split(token.Token, ":")[1])

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	resp.Body.Close()

	if resp.StatusCode == 500 {
		t.Fatal("unexpected 500 error")
	}

	return resp.StatusCode != http.StatusUnauthorized
}

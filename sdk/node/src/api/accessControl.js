const shared = require('../shared')

/**
 * Grants provide a mapping from guard objects (access tokens or X509
 * certificates) to a list of predefined Chain Core access policies.
 *
 * * **client-readwrite**: full access to the Client API
 * * **client-readonly**: access to read-only Client endpoints
 * * **network**: access to the Network API
 * * **monitoring**: access to monitoring-specific endpoints
 * * **internal**: access to multi-process synchronization endpoints (Raft, etc.)
 *
 * More info: {@link https://chain.com/docs/core/learn-more/authentication}
 * @typedef {Object} Grant
 * @global
 *
 * @property {String} guard_type
 * Type of credential, either 'access_token' or 'x509'.
 *
 * @property {Object} guard_data
 * Object containing data needed to identify the incoming credential.
 *
 * @property {String} policy
 * Authorization single polciy to attach to specific grant.
 */

/**
 * API for interacting with {@link Grant access grants}.
 *
 * More info: {@link https://chain.com/docs/core/learn-more/authentication}
 * @module AccessControlApi
 */
const accessControl = (client) => ({
  /**
   * Create a new access grant.
   *
   * @param {Object} params - Parameters for access grant creation.
   * @param {String} params.guard_type - Type of credential to guard with, either 'access_token' or 'x509'.
   * @param {Object} params.guard_data - Object containing data needed to identify the incoming credential.
   * @param {String} params.policy - Authorization polciy to attach to specific grant. See {@link Grant} for a list of available policiies.
   * @param {objectCallback} [callback] - Optional callback. Use instead of Promise return value as desired.
   * @returns {Promise} Status of created object.
   */
  create: (params , cb) =>
    shared.create(client, '/create-acl-grant', params, {skipArray: true}, cb),

  /**
   * Delete the specfiied access grant.
   *
   * @param {Object} params - Parameters for access grant deletion.
   * @param {String} params.guard_type - Type of credential to delete, either 'access_token' or 'x509'.
   * @param {Object} params.guard_data - Object containing data needed to identify the credential to be removed.
   * @param {String} params.policy - Authorization policy to remove. See {@link Grant} for a list of available policiies.
   * @param {objectCallback} [callback] - Optional callback. Use instead of Promise return value as desired.
   * @returns {Promise} Status of deleted object.
   */
  delete: (params, cb) => shared.tryCallback(
    client.request('/revoke-acl-grant', params),
    cb
  ),

  /**
   * Get all access grants.
   *
   * @param {pageCallback} [callback] - Optional callback. Use instead of Promise return value as desired.
   * @returns {Promise<Array<Grant>>} Requested page of results.
   */
  query: (cb) =>
    shared.query(client, 'accessTokens', '/list-acl-grants', {}, {cb}),
})

module.exports = accessControl

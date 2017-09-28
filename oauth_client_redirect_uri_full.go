/* 
 * Phone.com API
 *
 * This is a Phone.com api Swagger definition
 *
 * OpenAPI spec version: 1.0.0
 * Contact: apisupport@phone.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

// An OAuth Client Full Object is identical as a Summary Object.
type OauthClientRedirectUriFull struct {

	// OAuth Client Redirect URI ID
	Id int32 `json:"id,omitempty"`

	// Client details
	Client OauthClientFull `json:"client,omitempty"`

	// Redirect URI details
	RedirectUri RedirectUriFull `json:"redirect_uri,omitempty"`
}

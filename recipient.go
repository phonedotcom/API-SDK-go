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

type Recipient struct {

	// Phone number that will receive the SMS message
	Number string `json:"number,omitempty"`

	// Indicate the status of your SMS object. May be 'sent', 'received', 'queued', 'new' ...
	Status string `json:"status,omitempty"`
}

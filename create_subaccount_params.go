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

type CreateSubaccountParams struct {

	// Sub account password
	Username string `json:"username"`

	// Sub account password
	Password string `json:"password"`

	// Contact Object. See below for details.
	Contact ContactResponse `json:"contact,omitempty"`

	// Contact Object for billing purposes. See below for details.
	BillingContact ContactResponse `json:"billing_contact,omitempty"`
}

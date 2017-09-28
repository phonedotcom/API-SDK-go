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

type CreateExtensionParams struct {

	Voicemail VoicemailInput `json:"voicemail,omitempty"`

	CallNotifications CallNotifications `json:"call_notifications,omitempty"`

	// Caller ID
	CallerId string `json:"caller_id,omitempty"`

	// Extension type
	UsageType string `json:"usage_type,omitempty"`

	// Extension number (auto-generated if empty)
	Extension int32 `json:"extension,omitempty"`

	// Include in dial-by-name directory
	IncludeInDirectory string `json:"include_in_directory,omitempty"`

	// Name (auto-generated if empty)
	Name string `json:"name,omitempty"`

	// Contact name
	FullName string `json:"full_name,omitempty"`

	// Timezone
	Timezone string `json:"timezone,omitempty"`

	// Recording lookup object
	NameGreeting interface{} `json:"name_greeting,omitempty"`

	// Local area code
	LocalAreaCode string `json:"local_area_code,omitempty"`

	// Enable outgoing calls
	EnableOutboundCalls string `json:"enable_outbound_calls,omitempty"`

	// Enable Call Waiting
	EnableCallWaiting string `json:"enable_call_waiting,omitempty"`
}

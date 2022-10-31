// Code generated by goa v3.10.2, DO NOT EDIT.
//
// HTTP request path constructors for the admin service.
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package server

// UpdateAgentAdminPath returns the URL path to the admin service UpdateAgent HTTP endpoint.
func UpdateAgentAdminPath() string {
	return "/system/user/agent"
}

// RefreshConfigAdminPath returns the URL path to the admin service RefreshConfig HTTP endpoint.
func RefreshConfigAdminPath() string {
	return "/system/config/refresh"
}

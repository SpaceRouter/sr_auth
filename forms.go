package sr_auth

type userRolesResponse struct {
	Message string
	Ok      bool
	Role    Role
}

type userInfoResponse struct {
	Message  string
	Ok       bool
	UserInfo UserInfo
}

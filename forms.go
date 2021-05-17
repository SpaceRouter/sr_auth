package sr_auth

type userRolesResponse struct {
	Message string
	Ok      bool
	Roles   []Role
}

type userInfoResponse struct {
	Message  string
	Ok       bool
	UserInfo UserInfo
}

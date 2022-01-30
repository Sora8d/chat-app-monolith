package common

type MultipleUuids []string

type AuthInfoJTW struct {
	AccessToken  string
	RefreshToken string
}

type AuthInfo struct {
	Uuid string
}

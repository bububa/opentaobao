package charity

// QueryThirdUserHasAuthHsfRequest 结构体
type QueryThirdUserHasAuthHsfRequest struct {
	// appkey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 用户唯一标识
	UserKey string `json:"user_key,omitempty" xml:"user_key,omitempty"`
}

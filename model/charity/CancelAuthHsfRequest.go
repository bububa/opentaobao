package charity

// CancelAuthHsfRequest 结构体
type CancelAuthHsfRequest struct {
	// appkey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 用户唯一标识
	UserKey string `json:"user_key,omitempty" xml:"user_key,omitempty"`
}

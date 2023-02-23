package charity

// QueryCharityTimeTopApiHsfRequest 结构体
type QueryCharityTimeTopApiHsfRequest struct {
	// 第三方userKey 必传
	UserKey string `json:"user_key,omitempty" xml:"user_key,omitempty"`
	// appkey 必传
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 活动id 必传
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}

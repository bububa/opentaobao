package cloudgame

// AlibabaCgameAvatarUserbodyQueryResult 结构体
type AlibabaCgameAvatarUserbodyQueryResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应数据
	Data *AlibabaCgameAvatarUserbodyQueryData `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

package user

// UserSubscribe 结构体
type UserSubscribe struct {
	// 订购状况。应用订购者：subscribeUser;尚未订购：unsubscribeUser；非法用户：invalidateUser
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 订购开始时间。格式:yyyy-MM-dd HH:mm:ss
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 订购结束时间。格式:yyyy-MM-dd HH:mm:ss
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 0-无版本信息；1-初级版；2-中级版；3-高级版
	VersionNo int64 `json:"version_no,omitempty" xml:"version_no,omitempty"`
}

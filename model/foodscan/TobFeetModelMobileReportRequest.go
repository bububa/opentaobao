package foodscan

// TobFeetModelMobileReportRequest 结构体
type TobFeetModelMobileReportRequest struct {
	// 2020-05-20 00:00:00
	RelationType int64 `json:"relation_type,omitempty" xml:"relation_type,omitempty"`
	// 1男2女
	Gender int64 `json:"gender,omitempty" xml:"gender,omitempty"`
	// yyyy-MM-dd HH:mm:ss
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// yyyy-MM-dd HH:mm:ss
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 第几页
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
	// 用户唯一标识，可以是淘宝用户ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 淘宝昵称
	NickName string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
}

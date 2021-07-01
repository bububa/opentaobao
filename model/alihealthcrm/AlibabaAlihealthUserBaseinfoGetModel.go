package alihealthcrm

// AlibabaAlihealthUserBaseinfoGetModel 结构体
type AlibabaAlihealthUserBaseinfoGetModel struct {
	// 运动状态
	SportsStatus int64 `json:"sports_status,omitempty" xml:"sports_status,omitempty"`
	// 生日
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 体重
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 身高
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 性别
	Gender int64 `json:"gender,omitempty" xml:"gender,omitempty"`
	// 用户id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 用户头像信息
	HeadUrl string `json:"head_url,omitempty" xml:"head_url,omitempty"`
	// 养生目标
	HealthTarget string `json:"health_target,omitempty" xml:"health_target,omitempty"`
	// 体质状况
	BodyStatus string `json:"body_status,omitempty" xml:"body_status,omitempty"`
	// 健康问答
	HealthQa string `json:"health_qa,omitempty" xml:"health_qa,omitempty"`
	// 昵称
	NickName string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// 加密手机号串
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

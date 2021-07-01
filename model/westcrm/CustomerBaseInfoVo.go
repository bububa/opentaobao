package westcrm

// CustomerBaseInfoVo 结构体
type CustomerBaseInfoVo struct {
	// 积分
	Point int64 `json:"point,omitempty" xml:"point,omitempty"`
	// 头像
	Avatar string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 等级序号
	GradeNum int64 `json:"grade_num,omitempty" xml:"grade_num,omitempty"`
	// 园区id
	CampusId int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
	// 电话
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 性别
	Gender int64 `json:"gender,omitempty" xml:"gender,omitempty"`
	// 等级id
	GradeId int64 `json:"grade_id,omitempty" xml:"grade_id,omitempty"`
	// 用户id
	IbUserId int64 `json:"ib_user_id,omitempty" xml:"ib_user_id,omitempty"`
	// 等级名称
	GradeName string `json:"grade_name,omitempty" xml:"grade_name,omitempty"`
}

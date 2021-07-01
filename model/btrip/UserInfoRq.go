package btrip

// UserInfoRq 结构体
type UserInfoRq struct {
	// 第三方用户ID（注册签约时必填）
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 职务
	Position string `json:"position,omitempty" xml:"position,omitempty"`
	// 用户所在部门ID（注册签约时必填）
	DepartId int64 `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// 用户真实姓名（注册签约时必填）
	RealName string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	// 英文姓名请用"/"分隔，中间不能含有空格。顺序：姓/名(last_name/first_name))
	RealNameEn string `json:"real_name_en,omitempty" xml:"real_name_en,omitempty"`
	// 联系电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 工号
	JobNo string `json:"job_no,omitempty" xml:"job_no,omitempty"`
	// email
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 职务等级
	PositionLevel string `json:"position_level,omitempty" xml:"position_level,omitempty"`
}

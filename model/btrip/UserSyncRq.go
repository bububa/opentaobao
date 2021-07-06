package btrip

// UserSyncRq 结构体
type UserSyncRq struct {
	// 部门列表，depart_id | third_depart_id | third_depart_id_list只传其一，优先级为third_depart_id_list > third_depart_id > depart_id
	ThirdDepartIdList []string `json:"third_depart_id_list,omitempty" xml:"third_depart_id_list>string,omitempty"`
	// 第三方人员ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 姓名
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 手机号（除政府企业、特殊企业外，均为必填;如有特殊需求，请先联系商旅）
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 职级
	PositionLevel string `json:"position_level,omitempty" xml:"position_level,omitempty"`
	// 职位
	Position string `json:"position,omitempty" xml:"position,omitempty"`
	// 英文姓名请用"/"分隔，中间不能含有空格。顺序：姓/名(last_name/first_name))
	RealNameEn string `json:"real_name_en,omitempty" xml:"real_name_en,omitempty"`
	// 工号
	JobNo string `json:"job_no,omitempty" xml:"job_no,omitempty"`
	// 第三方部门ID
	ThirdDepartId string `json:"third_depart_id,omitempty" xml:"third_depart_id,omitempty"`
	// 商旅部门ID
	DepartId int64 `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// 是否离职（0 否 1是）
	LeaveStatus int64 `json:"leave_status,omitempty" xml:"leave_status,omitempty"`
}

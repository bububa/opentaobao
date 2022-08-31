package btrip

// OpenEmployeeInfo 结构体
type OpenEmployeeInfo struct {
	// 员工昵称。
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 第三方员工ID。
	ThirdPartUserId string `json:"third_part_user_id,omitempty" xml:"third_part_user_id,omitempty"`
	// 第三方员工工号。
	ThirdPartJobNo string `json:"third_part_job_no,omitempty" xml:"third_part_job_no,omitempty"`
}

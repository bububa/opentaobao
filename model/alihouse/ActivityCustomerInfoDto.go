package alihouse

// ActivityCustomerInfoDto 结构体
type ActivityCustomerInfoDto struct {
	// 团队名称
	TeamName string `json:"team_name,omitempty" xml:"team_name,omitempty"`
	// 顾问名称
	EmployeeName string `json:"employee_name,omitempty" xml:"employee_name,omitempty"`
	// 客户身份证
	CustomerCertNo string `json:"customer_cert_no,omitempty" xml:"customer_cert_no,omitempty"`
	// 外部顾问员工id
	OuterEmployeeId string `json:"outer_employee_id,omitempty" xml:"outer_employee_id,omitempty"`
	// 客户姓名
	CustomerMobile string `json:"customer_mobile,omitempty" xml:"customer_mobile,omitempty"`
	// 客户手机号
	CustomerName string `json:"customer_name,omitempty" xml:"customer_name,omitempty"`
	// 外部客户id
	OuterCustomerId string `json:"outer_customer_id,omitempty" xml:"outer_customer_id,omitempty"`
	// 是否删除
	IsDeleted int64 `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
}

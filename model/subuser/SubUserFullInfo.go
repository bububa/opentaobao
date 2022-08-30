package subuser

// SubUserFullInfo 结构体
type SubUserFullInfo struct {
	// 工作地点
	WorkLocation string `json:"work_location,omitempty" xml:"work_location,omitempty"`
	// 员工花名
	EmployeeNickname string `json:"employee_nickname,omitempty" xml:"employee_nickname,omitempty"`
	// 主账号企业邮箱
	UserEmail string `json:"user_email,omitempty" xml:"user_email,omitempty"`
	// 职务名称
	DutyName string `json:"duty_name,omitempty" xml:"duty_name,omitempty"`
	// 员工姓名
	EmployeeName string `json:"employee_name,omitempty" xml:"employee_name,omitempty"`
	// 入职员工工号
	EmployeeNum string `json:"employee_num,omitempty" xml:"employee_num,omitempty"`
	// 员工入职时间
	EntryDate string `json:"entry_date,omitempty" xml:"entry_date,omitempty"`
	// 部门名称
	DepartmentName string `json:"department_name,omitempty" xml:"department_name,omitempty"`
	// 子账号用户名
	SubNick string `json:"sub_nick,omitempty" xml:"sub_nick,omitempty"`
	// 主账号用户名
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 办公电话
	OfficePhone string `json:"office_phone,omitempty" xml:"office_phone,omitempty"`
	// 子账号企业邮箱
	SubuserEmail string `json:"subuser_email,omitempty" xml:"subuser_email,omitempty"`
	// 员工性别  1:男;  2:女
	Sex int64 `json:"sex,omitempty" xml:"sex,omitempty"`
	// 子账号Id
	SubId int64 `json:"sub_id,omitempty" xml:"sub_id,omitempty"`
	// 子账号当前状态：1正常，2卖家停用，3处罚冻结
	SubStatus int64 `json:"sub_status,omitempty" xml:"sub_status,omitempty"`
	// 部门Id
	DepartmentId int64 `json:"department_id,omitempty" xml:"department_id,omitempty"`
	// 职务等级
	DutyLevel int64 `json:"duty_level,omitempty" xml:"duty_level,omitempty"`
	// 直接上级的员工ID
	LeaderId int64 `json:"leader_id,omitempty" xml:"leader_id,omitempty"`
	// 父部门Id
	ParentDepartment int64 `json:"parent_department,omitempty" xml:"parent_department,omitempty"`
	// 职务Id
	DutyId int64 `json:"duty_id,omitempty" xml:"duty_id,omitempty"`
	// 主账号Id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 员工ID
	EmployeeId int64 `json:"employee_id,omitempty" xml:"employee_id,omitempty"`
	// 子账号是否参与分流 true:参与分流 false:未参与分流
	SubDispatchStatus bool `json:"sub_dispatch_status,omitempty" xml:"sub_dispatch_status,omitempty"`
	// 子账号是否已欠费 true:已欠费 false:未欠费
	SubOwedStatus bool `json:"sub_owed_status,omitempty" xml:"sub_owed_status,omitempty"`
}

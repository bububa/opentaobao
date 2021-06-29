package eleenterpriseemployee

// EmployeeInfoDto 
type EmployeeInfoDto struct {
    // 部门
    DeptName   string `json:"dept_name,omitempty" xml:"dept_name,omitempty"`
    // 手机号
    PhoneNumber   string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
    // 成本中心
    CostCenter   *CostCenter `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
    // 姓名
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 工号
    EmployeeNo   string `json:"employee_no,omitempty" xml:"employee_no,omitempty"`
}

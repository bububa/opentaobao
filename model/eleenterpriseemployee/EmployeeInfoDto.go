package eleenterpriseemployee

import (
	"sync"
)

// EmployeeInfoDto 结构体
type EmployeeInfoDto struct {
	// 部门
	DeptName string `json:"dept_name,omitempty" xml:"dept_name,omitempty"`
	// 手机号
	PhoneNumber string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 工号
	EmployeeNo string `json:"employee_no,omitempty" xml:"employee_no,omitempty"`
	// 成本中心
	CostCenter *CostCenter `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
}

var poolEmployeeInfoDto = sync.Pool{
	New: func() any {
		return new(EmployeeInfoDto)
	},
}

// GetEmployeeInfoDto() 从对象池中获取EmployeeInfoDto
func GetEmployeeInfoDto() *EmployeeInfoDto {
	return poolEmployeeInfoDto.Get().(*EmployeeInfoDto)
}

// ReleaseEmployeeInfoDto 释放EmployeeInfoDto
func ReleaseEmployeeInfoDto(v *EmployeeInfoDto) {
	v.DeptName = ""
	v.PhoneNumber = ""
	v.Name = ""
	v.EmployeeNo = ""
	v.CostCenter = nil
	poolEmployeeInfoDto.Put(v)
}

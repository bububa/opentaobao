package subuser

import (
	"sync"
)

// Department 结构体
type Department struct {
	// 部门下关联的子账号id列表
	SubUserIds []int64 `json:"sub_user_ids,omitempty" xml:"sub_user_ids>int64,omitempty"`
	// 部门名称
	DepartmentName string `json:"department_name,omitempty" xml:"department_name,omitempty"`
	// 部门ID
	DepartmentId int64 `json:"department_id,omitempty" xml:"department_id,omitempty"`
	// 当前部门的父部门ID
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
}

var poolDepartment = sync.Pool{
	New: func() any {
		return new(Department)
	},
}

// GetDepartment() 从对象池中获取Department
func GetDepartment() *Department {
	return poolDepartment.Get().(*Department)
}

// ReleaseDepartment 释放Department
func ReleaseDepartment(v *Department) {
	v.SubUserIds = v.SubUserIds[:0]
	v.DepartmentName = ""
	v.DepartmentId = 0
	v.ParentId = 0
	poolDepartment.Put(v)
}

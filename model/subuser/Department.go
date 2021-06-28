package subuser

// Department 
/* model for simplify = false
type Department struct {

    // 部门ID
    
    DepartmentId   int64 `json:"department_id,omitempty"`
    

    // 当前部门的父部门ID
    
    ParentId   int64 `json:"parent_id,omitempty"`
    

    // 部门名称
    
    DepartmentName   string `json:"department_name,omitempty"`
    

    // 部门下关联的子账号id列表
    
    SubUserIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"sub_user_ids,omitempty"`
    

}
*/

// Department 
type Department struct {

    // 部门ID
    DepartmentId   int64 `json:"department_id,omitempty"`

    // 当前部门的父部门ID
    ParentId   int64 `json:"parent_id,omitempty"`

    // 部门名称
    DepartmentName   string `json:"department_name,omitempty"`

    // 部门下关联的子账号id列表
    SubUserIds   []int64 `json:"sub_user_ids,omitempty"`

}

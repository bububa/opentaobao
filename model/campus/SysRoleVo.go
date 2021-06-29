package campus

// SysRoleVo 
type SysRoleVo struct {
    // id
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
    // roleName
    RoleName   string `json:"role_name,omitempty" xml:"role_name,omitempty"`
    // 是否被用户选择
    Granted   string `json:"granted,omitempty" xml:"granted,omitempty"`
    // 园区名称
    DeptName   string `json:"dept_name,omitempty" xml:"dept_name,omitempty"`
}

package campus

// RoleQueryParam 
type RoleQueryParam struct {
    // 支持多应用
    AppKeys   []string `json:"app_keys,omitempty" xml:"app_keys>string,omitempty"`
    // 角色名称
    RoleName   string `json:"role_name,omitempty" xml:"role_name,omitempty"`
    // 园区
    DeptId   string `json:"dept_id,omitempty" xml:"dept_id,omitempty"`
}

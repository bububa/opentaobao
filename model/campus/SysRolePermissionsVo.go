package campus

// SysRolePermissionsVo 
type SysRolePermissionsVo struct {

    // id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 角色key
    
    RoleKey   string `json:"role_key,omitempty" xml:"role_key,omitempty"`
    

    // 租户
    
    Tenant   string `json:"tenant,omitempty" xml:"tenant,omitempty"`
    

    // appKey
    
    AppKey   string `json:"app_key,omitempty" xml:"app_key,omitempty"`
    

    // 1
    
    DeptId   string `json:"dept_id,omitempty" xml:"dept_id,omitempty"`
    

    // 园区名称
    
    DeptName   string `json:"dept_name,omitempty" xml:"dept_name,omitempty"`
    

    // 角色名称
    
    RoleName   string `json:"role_name,omitempty" xml:"role_name,omitempty"`
    

    // 角色类型
    
    RoleType   string `json:"role_type,omitempty" xml:"role_type,omitempty"`
    

    // 角色描述
    
    RoleDesc   string `json:"role_desc,omitempty" xml:"role_desc,omitempty"`
    

    // permissions
    
    Permissions   []TreeNodeDto `json:"permissions,omitempty" xml:"permissions,omitempty"`
    

}

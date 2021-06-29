package campus

// UserRolesDTO 
type UserRolesDTO struct {
    // 用户账号
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 用户名称
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
    // roleList
    RoleList   []SysRoleDTO `json:"role_list,omitempty" xml:"role_list>sys_role_dto,omitempty"`
}

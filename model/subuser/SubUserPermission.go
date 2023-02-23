package subuser

// SubUserPermission 结构体
type SubUserPermission struct {
	// 子账号被赋予的角色信息(Role)列表。列表中的角色对象只有role_id，role_name，permissions信息
	Roles []Role `json:"roles,omitempty" xml:"roles>role,omitempty"`
	// 子账号被直接赋予的权限点列表
	Permissions []Permission `json:"permissions,omitempty" xml:"permissions>permission,omitempty"`
}

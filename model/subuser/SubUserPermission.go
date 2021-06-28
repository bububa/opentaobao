package subuser

// SubUserPermission 
/* model for simplify = false
type SubUserPermission struct {

    // 子账号被直接赋予的权限点列表
    
    Permissions  struct {
        Permission  []Permission `json:"permission,omitempty"`
    } `json:"permissions,omitempty"`
    

    // 子账号被赋予的角色信息(Role)列表。列表中的角色对象只有role_id，role_name，permissions信息
    
    Roles  struct {
        Role  []Role `json:"role,omitempty"`
    } `json:"roles,omitempty"`
    

}
*/

// SubUserPermission 
type SubUserPermission struct {

    // 子账号被直接赋予的权限点列表
    Permissions   []Permission `json:"permissions,omitempty"`

    // 子账号被赋予的角色信息(Role)列表。列表中的角色对象只有role_id，role_name，permissions信息
    Roles   []Role `json:"roles,omitempty"`

}

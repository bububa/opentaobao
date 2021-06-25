package subuser

// Role 
type Role struct {

    // 角色id
    RoleId   int64 `json:"role_id,omitempty"`

    // 角色名
    RoleName   string `json:"role_name,omitempty"`

    // 角色描述
    Description   string `json:"description,omitempty"`

    // 卖家Id
    SellerId   int64 `json:"seller_id,omitempty"`

    // 创建时间
    CreateTime   string `json:"create_time,omitempty"`

    // 修改时间
    ModifiedTime   string `json:"modified_time,omitempty"`

    // 所拥有权限
    Permissions   []Permission `json:"permissions,omitempty"`

}

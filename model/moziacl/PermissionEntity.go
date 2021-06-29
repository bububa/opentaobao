package moziacl

// PermissionEntity 
type PermissionEntity struct {
    // 权限描述
    PermissionDescription   string `json:"permission_description,omitempty" xml:"permission_description,omitempty"`
    // 权限归属人userId列表
    PermissionOwnerIdList   []string `json:"permission_owner_id_list,omitempty" xml:"permission_owner_id_list>string,omitempty"`
    // 最大过期时间
    MaxExpireDate   string `json:"max_expire_date,omitempty" xml:"max_expire_date,omitempty"`
    // 风险等级
    RiskLevel   string `json:"risk_level,omitempty" xml:"risk_level,omitempty"`
    // 权限英文名
    PermissionTitleEN   string `json:"permission_title_e_n,omitempty" xml:"permission_title_e_n,omitempty"`
    // 权限中文名
    PermissionTitle   string `json:"permission_title,omitempty" xml:"permission_title,omitempty"`
    // 权限code
    PermissionName   string `json:"permission_name,omitempty" xml:"permission_name,omitempty"`
    // 状态
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 公开申请策略
    PublicAttri   string `json:"public_attri,omitempty" xml:"public_attri,omitempty"`
    // 创建人
    Creator   *BucUser `json:"creator,omitempty" xml:"creator,omitempty"`
    // 创建时间
    CreatTime   string `json:"creat_time,omitempty" xml:"creat_time,omitempty"`
    // 创建时间
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    // 回收规则
    RevokeRule   string `json:"revoke_rule,omitempty" xml:"revoke_rule,omitempty"`
    // 是否可用
    IsActive   bool `json:"is_active,omitempty" xml:"is_active,omitempty"`
}

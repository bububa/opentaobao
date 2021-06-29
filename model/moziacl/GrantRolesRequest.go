package moziacl

// GrantRolesRequest 
type GrantRolesRequest struct {
    // 调用接口授角色的原因
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    // 目标应用名
    TargetAppName   string `json:"target_app_name,omitempty" xml:"target_app_name,omitempty"`
    // 扩展参数
    RequestMetaData   string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
    // 授角色主体
    Principal   *BucUserPrincipalParam `json:"principal,omitempty" xml:"principal,omitempty"`
    // 授予角色的过期时间
    ExpireDate   string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
    // 授予的角色的code列表
    RoleNames   []string `json:"role_names,omitempty" xml:"role_names>string,omitempty"`
}

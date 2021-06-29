package moziacl

// RemovePermissionsFromRoleRequest 
type RemovePermissionsFromRoleRequest struct {
    // 操作主体
    PrincipalParam   *BucUserPrincipalParam `json:"principal_param,omitempty" xml:"principal_param,omitempty"`
    // 角色所在应用name(应用唯一code)
    TargetAppName   string `json:"target_app_name,omitempty" xml:"target_app_name,omitempty"`
    // 角色name(角色唯一code，在ACL中全局唯一)
    RoleName   string `json:"role_name,omitempty" xml:"role_name,omitempty"`
    // 要移除的权限name列表(功能权限唯一code，在ACL中全局唯一)
    PermissionNames   []string `json:"permission_names,omitempty" xml:"permission_names>string,omitempty"`
    // 请求参数扩展字段
    RequestMetaData   string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
}

package moziacl

// RevokePermissionsRequest 
type RevokePermissionsRequest struct {
    // 回收主体对象
    Principal   *BucUserPrincipalParam `json:"principal,omitempty" xml:"principal,omitempty"`
    // 回收权限所在的app name
    TargetAppName   string `json:"target_app_name,omitempty" xml:"target_app_name,omitempty"`
    // 回收权限的name列表
    PermissionNames   []string `json:"permission_names,omitempty" xml:"permission_names>string,omitempty"`
    // 请求扩展字段
    RequestMetaData   string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
}

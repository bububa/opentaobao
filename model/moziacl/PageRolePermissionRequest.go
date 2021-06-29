package moziacl

// PageRolePermissionRequest 
type PageRolePermissionRequest struct {
    // 是否返回数据总数量
    ReturnTotalSize   bool `json:"return_total_size,omitempty" xml:"return_total_size,omitempty"`
    // 操作主体
    PrincipalParam   *BucUserPrincipalParam `json:"principal_param,omitempty" xml:"principal_param,omitempty"`
    // 角色所在的应用app name
    TargetAppName   string `json:"target_app_name,omitempty" xml:"target_app_name,omitempty"`
    // 每页返回数量
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 请求扩展字段
    RequestMetaData   string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
    // 角色下权限模糊匹配，如果传了，则将按照此模糊字段匹配角色下的权限
    FuzzyName   string `json:"fuzzy_name,omitempty" xml:"fuzzy_name,omitempty"`
    // 查询第几页
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    // 要查询的角色name
    RoleName   string `json:"role_name,omitempty" xml:"role_name,omitempty"`
}

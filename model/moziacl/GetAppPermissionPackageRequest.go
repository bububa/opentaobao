package moziacl

// GetAppPermissionPackageRequest 
type GetAppPermissionPackageRequest struct {
    // 是否查询权限套餐被开通的租户列表(如无必要，建议不要设置true，会增加额外查询)
    ShowRealmInfo   bool `json:"show_realm_info,omitempty" xml:"show_realm_info,omitempty"`
    // 是否返回数据总量
    ReturnTotalSize   bool `json:"return_total_size,omitempty" xml:"return_total_size,omitempty"`
    // 操作主体
    PrincipalParam   *BucUserPrincipalParam `json:"principal_param,omitempty" xml:"principal_param,omitempty"`
    // 要查询的应用的appname
    TargetAppName   string `json:"target_app_name,omitempty" xml:"target_app_name,omitempty"`
    // 每页数据条数
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 是否查询权限套餐中包含的角色、权限、数据权限 的数量。(如无必要，建议不要设置为true，会增加额外查询)
    ShowELementCount   bool `json:"show_e_lement_count,omitempty" xml:"show_e_lement_count,omitempty"`
    // 请求扩展字段
    RequestMetaData   string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
    // 权限套餐名模糊匹配
    FuzzyName   string `json:"fuzzy_name,omitempty" xml:"fuzzy_name,omitempty"`
    // 查询第几页
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
}

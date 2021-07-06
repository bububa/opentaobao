package moziacl

// UpdateRolesToPermissionPackageRequest 结构体
type UpdateRolesToPermissionPackageRequest struct {
	// 角色唯一标识列表
	RoleNames []string `json:"role_names,omitempty" xml:"role_names>string,omitempty"`
	// 权限套餐唯一标识
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 请求扩展字段
	RequestMetaData string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
	// 请求主体
	PrincipalParam *BucUserPrincipalParam `json:"principal_param,omitempty" xml:"principal_param,omitempty"`
}

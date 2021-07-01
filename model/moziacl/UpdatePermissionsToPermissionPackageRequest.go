package moziacl

// UpdatePermissionsToPermissionPackageRequest 结构体
type UpdatePermissionsToPermissionPackageRequest struct {
	// 请求主体
	PrincipalParam *BucUserPrincipalParam `json:"principal_param,omitempty" xml:"principal_param,omitempty"`
	// 权限套餐唯一标识
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 权限唯一标识列表
	PermissionNames []string `json:"permission_names,omitempty" xml:"permission_names>string,omitempty"`
	// 请求扩展字段，返回数据中扩展字段值与此相同
	RequestMetaData string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
}

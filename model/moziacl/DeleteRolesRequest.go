package moziacl

// DeleteRolesRequest 结构体
type DeleteRolesRequest struct {
	// 要删除的角色code列表
	Names []string `json:"names,omitempty" xml:"names>string,omitempty"`
	// 操作主体
	PrincipalParam *BucUserPrincipalParam `json:"principal_param,omitempty" xml:"principal_param,omitempty"`
}

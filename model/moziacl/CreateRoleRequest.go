package moziacl

// CreateRoleRequest 结构体
type CreateRoleRequest struct {
	// 角色包含的权限name列表
	AddPermissionNames []string `json:"add_permission_names,omitempty" xml:"add_permission_names>string,omitempty"`
	// 角色审批人userId列表
	ApproverUserIds []string `json:"approver_user_ids,omitempty" xml:"approver_user_ids>string,omitempty"`
	// 角色归属的应用name，不传则以appKey对应的应用为准
	TargetAppName string `json:"target_app_name,omitempty" xml:"target_app_name,omitempty"`
	// 角色描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 角色中文名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 角色类型，重点注意：如果创建租户内角色，则必须填“Realm_Role”，ISV开发者创建应用侧角色，则可不填
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 请求扩展字段
	RequestMetaData string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
	// 角色包含的数据权限，以json传入
	DataPermissionJsonStr string `json:"data_permission_json_str,omitempty" xml:"data_permission_json_str,omitempty"`
	// 角色英文名称
	TitleEN string `json:"title_e_n,omitempty" xml:"title_e_n,omitempty"`
	// 角色审批规则类型(random、self 两种类型)
	RuleType string `json:"rule_type,omitempty" xml:"rule_type,omitempty"`
	// 角色申请规则(1表示公开可申请，2表示公开不可申请、3表示不公开)
	PublicAttri string `json:"public_attri,omitempty" xml:"public_attri,omitempty"`
	// 风险等级：L表示低风险，M表示中风，H表示高风险
	AssignLevel string `json:"assign_level,omitempty" xml:"assign_level,omitempty"`
	// 角色唯一code，在ACL中全局唯一
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 角色回收策略(REVOKE代表用户所在部门发生变化时，该用户的权限将被回收，RESERVE代表用户所在部门发生变化时，该用户的权限将被保留，TRANSFER_REVOKE代表个人主动转岗时回收 - 只在小二主动申请转岗时回收本权限)
	RevokeRule string `json:"revoke_rule,omitempty" xml:"revoke_rule,omitempty"`
	// 扩展字段
	ExtentionMap string `json:"extention_map,omitempty" xml:"extention_map,omitempty"`
	// 操作主体
	PrincipalParam *BucUserPrincipalParam `json:"principal_param,omitempty" xml:"principal_param,omitempty"`
	// 操作人的userId
	OperatorUserId int64 `json:"operator_user_id,omitempty" xml:"operator_user_id,omitempty"`
	// 角色归属人userId
	OwnerUserId int64 `json:"owner_user_id,omitempty" xml:"owner_user_id,omitempty"`
	// 是否数据权限，角色没有挂载数据权限，则为false
	IsData bool `json:"is_data,omitempty" xml:"is_data,omitempty"`
}

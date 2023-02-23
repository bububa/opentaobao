package alihouse

// UserPermissionInfoDto 结构体
type UserPermissionInfoDto struct {
	// 装修场景主体
	SceneInfoList []AstoreSceneInfoDto `json:"scene_info_list,omitempty" xml:"scene_info_list>astore_scene_info_dto,omitempty"`
	// 权限类型
	FunctionList []string `json:"function_list,omitempty" xml:"function_list>string,omitempty"`
	// 1-公司  2-城市品牌店 3-品牌  4-标准门店  5-TP运营账号 6-经纪人/置业顾问
	OuterTargetId string `json:"outer_target_id,omitempty" xml:"outer_target_id,omitempty"`
	// 外部商家id
	OuterSellerId string `json:"outer_seller_id,omitempty" xml:"outer_seller_id,omitempty"`
	// 业务域
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 同步的数据类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 外部门店ID
	OuterTargetType int64 `json:"outer_target_type,omitempty" xml:"outer_target_type,omitempty"`
}

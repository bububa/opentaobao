package alihouse

import (
	"sync"
)

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

var poolUserPermissionInfoDto = sync.Pool{
	New: func() any {
		return new(UserPermissionInfoDto)
	},
}

// GetUserPermissionInfoDto() 从对象池中获取UserPermissionInfoDto
func GetUserPermissionInfoDto() *UserPermissionInfoDto {
	return poolUserPermissionInfoDto.Get().(*UserPermissionInfoDto)
}

// ReleaseUserPermissionInfoDto 释放UserPermissionInfoDto
func ReleaseUserPermissionInfoDto(v *UserPermissionInfoDto) {
	v.SceneInfoList = v.SceneInfoList[:0]
	v.FunctionList = v.FunctionList[:0]
	v.OuterTargetId = ""
	v.OuterSellerId = ""
	v.BusinessType = 0
	v.Type = 0
	v.OuterTargetType = 0
	poolUserPermissionInfoDto.Put(v)
}

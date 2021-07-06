package westcrm

// ShopListVo 结构体
type ShopListVo struct {
	// 商铺类型
	ShopType string `json:"shop_type,omitempty" xml:"shop_type,omitempty"`
	// 商铺ID
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 商圈id
	CampusId string `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
	// 一级业态
	TypeLevel1 string `json:"type_level1,omitempty" xml:"type_level1,omitempty"`
	// 二级业态
	TypeLevel2 string `json:"type_level2,omitempty" xml:"type_level2,omitempty"`
	// 目标客群结构体
	Tags *ThirdTagsVo `json:"tags,omitempty" xml:"tags,omitempty"`
}

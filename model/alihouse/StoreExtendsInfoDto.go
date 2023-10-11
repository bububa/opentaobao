package alihouse

// StoreExtendsInfoDto 结构体
type StoreExtendsInfoDto struct {
	// 外部门店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 扩展信息
	ExtendsInfo string `json:"extends_info,omitempty" xml:"extends_info,omitempty"`
}

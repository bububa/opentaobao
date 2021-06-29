package ascpchannel

// Packageitemdtos 
type Packageitemdtos struct {
    // 包裹号
    PackageId   int64 `json:"package_id,omitempty" xml:"package_id,omitempty"`
    // 货品id
    ScItemId   int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
    // 商品名称
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    // 商品数量
    ItemQuantity   int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
    // 包裹明细签收信息列表
    PackageItemSignInfoDtoList   []Packageitemsigninfodtolist `json:"package_item_sign_info_dto_list,omitempty" xml:"package_item_sign_info_dto_list>packageitemsigninfodtolist,omitempty"`
}

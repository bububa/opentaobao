package alsc

// ComboInfo 
type ComboInfo struct {

    // 套餐菜品做法明细
    CookingMethodsInfoList   []AttachInfo `json:"cooking_methods_info_list,omitempty"`

    // 套餐菜品配料明细
    IngredientsInfoList   []AttachInfo `json:"ingredients_info_list,omitempty"`

    // 商品数量
    ItemCount   int64 `json:"item_count,omitempty"`

    // 商品id
    OutItemId   string `json:"out_item_id,omitempty"`

    // 商品名称
    OutItemName   string `json:"out_item_name,omitempty"`

    // 商品规格id
    OutSkuId   string `json:"out_sku_id,omitempty"`

    // 商品规格名称
    OutSkuName   string `json:"out_sku_name,omitempty"`

    // 单价
    Price   int64 `json:"price,omitempty"`

    // 单位
    Unit   string `json:"unit,omitempty"`

    // 重量
    Weight   string `json:"weight,omitempty"`

}

package travel

// PontusTravelItemSkuInfo 
type PontusTravelItemSkuInfo struct {

    // 套餐名称
    PackageName   string `json:"package_name,omitempty"`

    // 映射到具体日历价格套餐的外部商家编码
    OuterSkuId   string `json:"outer_sku_id,omitempty"`

    // 套餐描述
    PackageDesc   string `json:"package_desc,omitempty"`

    // 套餐的日历价格库存。如果是预约商品，只需要填写一个Price，并且，不需要填写Price中的date字段不填，且预约商品只有成人价格和库存。
    Prices   []PontusTravelPrices `json:"prices,omitempty"`

    // 套餐操作仅限于删除[delete]和修改[modify]&quot;, 该操作修改的是套餐的名称和outer_sku_id
    PackageOperation   string `json:"package_operation,omitempty"`

    // package_operation 为midofy时，outer_sku_id将被该值覆盖
    UpdateOuterSkuId   string `json:"update_outer_sku_id,omitempty"`

    // 套餐关联的产品元素信息
    Products   []PontusTravelProduct `json:"products,omitempty"`

    // combos
    Combos   string `json:"combos,omitempty"`

    // 邮轮房型ID，新版本有值
    RoomTypeId   int64 `json:"room_type_id,omitempty"`

    // 邮轮房型类型
    RoomType   int64 `json:"room_type,omitempty"`

    // 邮轮房型名称
    RoomTypeName   string `json:"room_type_name,omitempty"`

    // 人数
    PeopleNumber   int64 `json:"people_number,omitempty"`

    // 下单人数是否与房型人数一致
    OrderCountMatch   bool `json:"order_count_match,omitempty"`

}

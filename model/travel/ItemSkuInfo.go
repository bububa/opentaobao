package travel

// ItemSkuInfo 
/* model for simplify = false
type ItemSkuInfo struct {

    // 套餐关联的产品元素信息
    
    Products  struct {
        Product  []Product `json:"product,omitempty"`
    } `json:"products,omitempty"`
    

    // 套餐id
    
    PackageId   int64 `json:"package_id,omitempty"`
    

    // 套餐下面对应商品元素信息 仅针对新版商品
    
    Combos   string `json:"combos,omitempty"`
    

    // 套餐的日历价格库存。如果是预约商品，只需要填写一个Price，并且，不需要填写Price中的date字段不填，且预约商品只有成人价格和库存。
    
    Prices  struct {
        Prices  []Prices `json:"prices,omitempty"`
    } `json:"prices,omitempty"`
    

    // 套餐名称
    
    PackageName   string `json:"package_name,omitempty"`
    

    // 套餐描述
    
    PackageDesc   string `json:"package_desc,omitempty"`
    

    // 映射到具体日历价格套餐的外部商家编码
    
    OuterSkuId   string `json:"outer_sku_id,omitempty"`
    

    // 邮轮房型ID
    
    RoomTypeId   int64 `json:"room_type_id,omitempty"`
    

    // 邮轮房型类型
    
    RoomType   int64 `json:"room_type,omitempty"`
    

    // 邮轮房型名称
    
    RoomTypeName   string `json:"room_type_name,omitempty"`
    

    // 邮轮房型人数
    
    PeopleNumber   int64 `json:"people_number,omitempty"`
    

    // 邮轮下单是否限制人数和房型人数一致
    
    OrderCountMatch   bool `json:"order_count_match,omitempty"`
    

    // 必填，套餐操作：1-增加一个套餐，2-删除一个套餐（根据outer_sku_id定位该套餐），3-覆盖修改一个套餐（根据outer_sku_id定位该套餐）
    
    PackageOperation   int64 `json:"package_operation,omitempty"`
    

}
*/

// ItemSkuInfo 
type ItemSkuInfo struct {

    // 套餐关联的产品元素信息
    Products   []Product `json:"products,omitempty"`

    // 套餐id
    PackageId   int64 `json:"package_id,omitempty"`

    // 套餐下面对应商品元素信息 仅针对新版商品
    Combos   string `json:"combos,omitempty"`

    // 套餐的日历价格库存。如果是预约商品，只需要填写一个Price，并且，不需要填写Price中的date字段不填，且预约商品只有成人价格和库存。
    Prices   []Prices `json:"prices,omitempty"`

    // 套餐名称
    PackageName   string `json:"package_name,omitempty"`

    // 套餐描述
    PackageDesc   string `json:"package_desc,omitempty"`

    // 映射到具体日历价格套餐的外部商家编码
    OuterSkuId   string `json:"outer_sku_id,omitempty"`

    // 邮轮房型ID
    RoomTypeId   int64 `json:"room_type_id,omitempty"`

    // 邮轮房型类型
    RoomType   int64 `json:"room_type,omitempty"`

    // 邮轮房型名称
    RoomTypeName   string `json:"room_type_name,omitempty"`

    // 邮轮房型人数
    PeopleNumber   int64 `json:"people_number,omitempty"`

    // 邮轮下单是否限制人数和房型人数一致
    OrderCountMatch   bool `json:"order_count_match,omitempty"`

    // 必填，套餐操作：1-增加一个套餐，2-删除一个套餐（根据outer_sku_id定位该套餐），3-覆盖修改一个套餐（根据outer_sku_id定位该套餐）
    PackageOperation   int64 `json:"package_operation,omitempty"`

}

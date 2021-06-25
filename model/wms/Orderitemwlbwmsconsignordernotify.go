package wms

// Orderitemwlbwmsconsignordernotify 
type Orderitemwlbwmsconsignordernotify struct {

    // ERP订单明细行号ID，数字类型
    OrderItemId   string `json:"order_item_id,omitempty"`

    // 货主ID
    OwnerUserId   string `json:"owner_user_id,omitempty"`

    // 货主名称
    OwnerUserName   string `json:"owner_user_name,omitempty"`

    // 店铺ID
    UserId   string `json:"user_id,omitempty"`

    // 店铺名称
    UserName   string `json:"user_name,omitempty"`

    // 平台子交易编码
    SubSourceCode   string `json:"sub_source_code,omitempty"`

    // 平台交易订单编码，如果不传入淘系交易订单，子订单系统自动标示此商品为赠品；
    OrderSourceCode   string `json:"order_source_code,omitempty"`

    // ERP商品ID
    ItemId   string `json:"item_id,omitempty"`

    // 商品名称
    ItemName   string `json:"item_name,omitempty"`

    // 商品数量
    ItemQuantity   int64 `json:"item_quantity,omitempty"`

    // 库存类型
    InventoryType   int64 `json:"inventory_type,omitempty"`

    // 销售价格
    ItemPrice   int64 `json:"item_price,omitempty"`

    // 商品优惠金额
    DiscountAmount   int64 `json:"discount_amount,omitempty"`

    // 商品成交价格=销售价格-优惠金额
    ActualPrice   int64 `json:"actual_price,omitempty"`

    // 订单商品拓展属性数据
    ExtendFields   string `json:"extend_fields,omitempty"`

    // ERP店铺编码
    ShopCode   string `json:"shop_code,omitempty"`

}

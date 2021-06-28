package travel

// FullTravelItem 
/* model for simplify = false
type FullTravelItem struct {

    // 商品基本信息
    
    BaseInfo  *struct {
        ItemBaseInfo  *ItemBaseInfo `json:"item_base_info,omitempty"`
    } `json:"base_info,omitempty"`
    

    // 预定规则
    
    BookingRules  struct {
        BookingRuleInfo  []BookingRuleInfo `json:"booking_rule_info,omitempty"`
    } `json:"booking_rules,omitempty"`
    

    // 创建时间
    
    Created   string `json:"created,omitempty"`
    

    // 自由行相关信息
    
    FreedomItemExt  *struct {
        FreedomItemExt  *FreedomItemExt `json:"freedom_item_ext,omitempty"`
    } `json:"freedom_item_ext,omitempty"`
    

    // 跟团游相关信息
    
    GroupItemExt  *struct {
        GroupItemExt  *GroupItemExt `json:"group_item_ext,omitempty"`
    } `json:"group_item_ext,omitempty"`
    

    // 商品类型
    
    ItemType   int64 `json:"item_type,omitempty"`
    

    // 商品状态。0,1正常;-1:用户删除;-2:用户下架;-3 小二下架;-4 小二删除;-5 从未上架;-9 被处罚
    
    ItemStatus   int64 `json:"item_status,omitempty"`
    

    // 商品id1
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // 行程信息
    
    Itineraries  struct {
        ItemItineraryInfo  []ItemItineraryInfo `json:"item_itinerary_info,omitempty"`
    } `json:"itineraries,omitempty"`
    

    // 修改时间
    
    Modified   string `json:"modified,omitempty"`
    

    // 退改规则信息
    
    RefundInfo  *struct {
        ItemRefundInfo  *ItemRefundInfo `json:"item_refund_info,omitempty"`
    } `json:"refund_info,omitempty"`
    

    // 销售属性信息
    
    SaleInfo  *struct {
        ItemSaleInfo  *ItemSaleInfo `json:"item_sale_info,omitempty"`
    } `json:"sale_info,omitempty"`
    

    // 卖家昵称
    
    SellerNick   string `json:"seller_nick,omitempty"`
    

    // 卖家id
    
    SellerId   int64 `json:"seller_id,omitempty"`
    

    // sku信息
    
    SkuInfos  struct {
        ItemSkuInfo  []ItemSkuInfo `json:"item_sku_info,omitempty"`
    } `json:"sku_infos,omitempty"`
    

    // 同城活动商品相关信息
    
    TcwlItemExt  *struct {
        TcwlItemExt  *TcwlItemExt `json:"tcwl_item_ext,omitempty"`
    } `json:"tcwl_item_ext,omitempty"`
    

    // 航旅度假TOP API3.0 邮轮扩展信息结构
    
    CruiseItemExt  *struct {
        CruiseItemExt  *CruiseItemExt `json:"cruise_item_ext,omitempty"`
    } `json:"cruise_item_ext,omitempty"`
    

    // 商品特征数据
    
    Features   string `json:"features,omitempty"`
    

    // 新版行程信息
    
    RefTrip   string `json:"ref_trip,omitempty"`
    

    // 产品亮点
    
    HighLights  struct {
        ProductHighLights  []ProductHighLights `json:"product_high_lights,omitempty"`
    } `json:"high_lights,omitempty"`
    

}
*/

// FullTravelItem 
type FullTravelItem struct {

    // 商品基本信息
    BaseInfo   *ItemBaseInfo `json:"base_info,omitempty"`

    // 预定规则
    BookingRules   []BookingRuleInfo `json:"booking_rules,omitempty"`

    // 创建时间
    Created   string `json:"created,omitempty"`

    // 自由行相关信息
    FreedomItemExt   *FreedomItemExt `json:"freedom_item_ext,omitempty"`

    // 跟团游相关信息
    GroupItemExt   *GroupItemExt `json:"group_item_ext,omitempty"`

    // 商品类型
    ItemType   int64 `json:"item_type,omitempty"`

    // 商品状态。0,1正常;-1:用户删除;-2:用户下架;-3 小二下架;-4 小二删除;-5 从未上架;-9 被处罚
    ItemStatus   int64 `json:"item_status,omitempty"`

    // 商品id1
    ItemId   int64 `json:"item_id,omitempty"`

    // 行程信息
    Itineraries   []ItemItineraryInfo `json:"itineraries,omitempty"`

    // 修改时间
    Modified   string `json:"modified,omitempty"`

    // 退改规则信息
    RefundInfo   *ItemRefundInfo `json:"refund_info,omitempty"`

    // 销售属性信息
    SaleInfo   *ItemSaleInfo `json:"sale_info,omitempty"`

    // 卖家昵称
    SellerNick   string `json:"seller_nick,omitempty"`

    // 卖家id
    SellerId   int64 `json:"seller_id,omitempty"`

    // sku信息
    SkuInfos   []ItemSkuInfo `json:"sku_infos,omitempty"`

    // 同城活动商品相关信息
    TcwlItemExt   *TcwlItemExt `json:"tcwl_item_ext,omitempty"`

    // 航旅度假TOP API3.0 邮轮扩展信息结构
    CruiseItemExt   *CruiseItemExt `json:"cruise_item_ext,omitempty"`

    // 商品特征数据
    Features   string `json:"features,omitempty"`

    // 新版行程信息
    RefTrip   string `json:"ref_trip,omitempty"`

    // 产品亮点
    HighLights   []ProductHighLights `json:"high_lights,omitempty"`

}

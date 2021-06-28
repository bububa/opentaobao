package travel

// PontusTravelFullTravelItem 
/* model for simplify = false
type PontusTravelFullTravelItem struct {

    // 商品类型
    
    ItemType   int64 `json:"item_type,omitempty"`
    

    // 商品状态
    
    ItemStatus   int64 `json:"item_status,omitempty"`
    

    // 行程信息
    
    Itineraries  struct {
        PontusTravelItemItineraryInfo  []PontusTravelItemItineraryInfo `json:"pontus_travel_item_itinerary_info,omitempty"`
    } `json:"itineraries,omitempty"`
    

    // 商品id
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // sku信息
    
    SkuInfos  struct {
        PontusTravelItemSkuInfo  []PontusTravelItemSkuInfo `json:"pontus_travel_item_sku_info,omitempty"`
    } `json:"sku_infos,omitempty"`
    

    // 销售属性信息
    
    SaleInfo  *struct {
        PontusTravelItemSaleInfo  *PontusTravelItemSaleInfo `json:"pontus_travel_item_sale_info,omitempty"`
    } `json:"sale_info,omitempty"`
    

    // 跟团游相关信息
    
    GroupItemExt  *struct {
        PontusTravelGroupItemExt  *PontusTravelGroupItemExt `json:"pontus_travel_group_item_ext,omitempty"`
    } `json:"group_item_ext,omitempty"`
    

    // 修改时间
    
    Modified   string `json:"modified,omitempty"`
    

    // 商品基本信息
    
    BaseInfo  *struct {
        PontusTravelItemBaseInfo  *PontusTravelItemBaseInfo `json:"pontus_travel_item_base_info,omitempty"`
    } `json:"base_info,omitempty"`
    

    // 预定规则
    
    BookingRules  struct {
        PontusTravelBookingRuleInfo  []PontusTravelBookingRuleInfo `json:"pontus_travel_booking_rule_info,omitempty"`
    } `json:"booking_rules,omitempty"`
    

    // 卖家昵称
    
    SellerNick   string `json:"seller_nick,omitempty"`
    

    // 退改规则信息
    
    RefundInfo  *struct {
        PontusTravelItemRefundInfo  *PontusTravelItemRefundInfo `json:"pontus_travel_item_refund_info,omitempty"`
    } `json:"refund_info,omitempty"`
    

    // 创建时间
    
    Created   string `json:"created,omitempty"`
    

    // 卖家id
    
    SellerId   int64 `json:"seller_id,omitempty"`
    

    // 自由行相关信息
    
    FreedomItemExt  *struct {
        PontusTravelFreedomItemExt  *PontusTravelFreedomItemExt `json:"pontus_travel_freedom_item_ext,omitempty"`
    } `json:"freedom_item_ext,omitempty"`
    

    // 商品特征数据
    
    Features   string `json:"features,omitempty"`
    

    // 产品亮点
    
    HighLights  struct {
        HighLights  []HighLights `json:"high_lights,omitempty"`
    } `json:"high_lights,omitempty"`
    

    // 同城活动商品相关信息
    
    TcwlItemExt   string `json:"tcwl_item_ext,omitempty"`
    

    // 航旅度假TOP API3.0 邮轮扩展信息结构
    
    CruiseItemExt   string `json:"cruise_item_ext,omitempty"`
    

    // 新版行程信息
    
    RefTrip   string `json:"ref_trip,omitempty"`
    

}
*/

// PontusTravelFullTravelItem 
type PontusTravelFullTravelItem struct {

    // 商品类型
    ItemType   int64 `json:"item_type,omitempty"`

    // 商品状态
    ItemStatus   int64 `json:"item_status,omitempty"`

    // 行程信息
    Itineraries   []PontusTravelItemItineraryInfo `json:"itineraries,omitempty"`

    // 商品id
    ItemId   int64 `json:"item_id,omitempty"`

    // sku信息
    SkuInfos   []PontusTravelItemSkuInfo `json:"sku_infos,omitempty"`

    // 销售属性信息
    SaleInfo   *PontusTravelItemSaleInfo `json:"sale_info,omitempty"`

    // 跟团游相关信息
    GroupItemExt   *PontusTravelGroupItemExt `json:"group_item_ext,omitempty"`

    // 修改时间
    Modified   string `json:"modified,omitempty"`

    // 商品基本信息
    BaseInfo   *PontusTravelItemBaseInfo `json:"base_info,omitempty"`

    // 预定规则
    BookingRules   []PontusTravelBookingRuleInfo `json:"booking_rules,omitempty"`

    // 卖家昵称
    SellerNick   string `json:"seller_nick,omitempty"`

    // 退改规则信息
    RefundInfo   *PontusTravelItemRefundInfo `json:"refund_info,omitempty"`

    // 创建时间
    Created   string `json:"created,omitempty"`

    // 卖家id
    SellerId   int64 `json:"seller_id,omitempty"`

    // 自由行相关信息
    FreedomItemExt   *PontusTravelFreedomItemExt `json:"freedom_item_ext,omitempty"`

    // 商品特征数据
    Features   string `json:"features,omitempty"`

    // 产品亮点
    HighLights   []HighLights `json:"high_lights,omitempty"`

    // 同城活动商品相关信息
    TcwlItemExt   string `json:"tcwl_item_ext,omitempty"`

    // 航旅度假TOP API3.0 邮轮扩展信息结构
    CruiseItemExt   string `json:"cruise_item_ext,omitempty"`

    // 新版行程信息
    RefTrip   string `json:"ref_trip,omitempty"`

}

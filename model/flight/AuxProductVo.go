package flight

// AuxProductVo 
/* model for simplify = false
type AuxProductVo struct {

    // 结构化的行李说明
    
    BaggageVo  *struct {
        BaggageVo  *BaggageVo `json:"baggage_vo,omitempty"`
    } `json:"baggage_vo,omitempty"`
    

    // 基准价
    
    BasePrice   int64 `json:"base_price,omitempty"`
    

    // 柜台价
    
    CounterPrice   int64 `json:"counter_price,omitempty"`
    

    // 销售价格
    
    OnlinePrice   int64 `json:"online_price,omitempty"`
    

    // 外部outerId
    
    OuterId   string `json:"outer_id,omitempty"`
    

    // 辅营产品名称
    
    ProductName   string `json:"product_name,omitempty"`
    

    // deprecated, 不建议使用
    
    ProductType   int64 `json:"product_type,omitempty"`
    

    // 销售类型 1:一次销售，2:二次销售，3:不限
    
    SaleType   int64 `json:"sale_type,omitempty"`
    

    // 选座规格信息
    
    SeatVo  *struct {
        SeatVo  *SeatVo `json:"seat_vo,omitempty"`
    } `json:"seat_vo,omitempty"`
    

    // 产品类型 4:付费行李，6:值机，7:餐食, 8:值机及选座
    
    UnityProductType   int64 `json:"unity_product_type,omitempty"`
    

}
*/

// AuxProductVo 
type AuxProductVo struct {

    // 结构化的行李说明
    BaggageVo   *BaggageVo `json:"baggage_vo,omitempty"`

    // 基准价
    BasePrice   int64 `json:"base_price,omitempty"`

    // 柜台价
    CounterPrice   int64 `json:"counter_price,omitempty"`

    // 销售价格
    OnlinePrice   int64 `json:"online_price,omitempty"`

    // 外部outerId
    OuterId   string `json:"outer_id,omitempty"`

    // 辅营产品名称
    ProductName   string `json:"product_name,omitempty"`

    // deprecated, 不建议使用
    ProductType   int64 `json:"product_type,omitempty"`

    // 销售类型 1:一次销售，2:二次销售，3:不限
    SaleType   int64 `json:"sale_type,omitempty"`

    // 选座规格信息
    SeatVo   *SeatVo `json:"seat_vo,omitempty"`

    // 产品类型 4:付费行李，6:值机，7:餐食, 8:值机及选座
    UnityProductType   int64 `json:"unity_product_type,omitempty"`

}

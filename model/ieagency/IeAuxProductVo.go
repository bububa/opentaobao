package ieagency

// IeAuxProductVo 
type IeAuxProductVo struct {

    // 产品外部ID
    
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    

    // 产品名称
    
    ProductName   string `json:"product_name,omitempty" xml:"product_name,omitempty"`
    

    // 产品类型 3:贵宾厅，4:行李，5:选座，6:套餐，7:CIP，11:值机，12:餐食, 13:值机及选座
    
    ProductType   int64 `json:"product_type,omitempty" xml:"product_type,omitempty"`
    

    // 销售类型 1:一次销售，2:二次销售，3:不限
    
    SaleType   int64 `json:"sale_type,omitempty" xml:"sale_type,omitempty"`
    

    // 商品销售价格单位人民币分，打包行李的单价
    
    OnlinePrice   int64 `json:"online_price,omitempty" xml:"online_price,omitempty"`
    

    // 行李产品规格信息(仅当productType=4有效）
    
    BaggageVo   *IeBaggageVo `json:"baggage_vo,omitempty" xml:"baggage_vo,omitempty"`
    

    // 选座产品信息(仅当productType=3或8有效）
    
    SeatVo   *IeSeatVo `json:"seat_vo,omitempty" xml:"seat_vo,omitempty"`
    

}

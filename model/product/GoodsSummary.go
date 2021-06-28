package product

// GoodsSummary 
/* model for simplify = false
type GoodsSummary struct {

    // max_price_cent 产品最高价格
    
    MaxPriceCent   int64 `json:"max_price_cent,omitempty"`
    

    // id 产品id
    
    Id   string `json:"id,omitempty"`
    

    // unit 产品单位
    
    Unit   string `json:"unit,omitempty"`
    

    // detail_url 产品详情url
    
    DetailUrl   string `json:"detail_url,omitempty"`
    

    // supplier_id 产品对应供应商id
    
    SupplierId   string `json:"supplier_id,omitempty"`
    

    // subject 产品标题
    
    Subject   string `json:"subject,omitempty"`
    

    // min_price_cent 产品最低价
    
    MinPriceCent   int64 `json:"min_price_cent,omitempty"`
    

    // thumb_url 产品缩略图
    
    ThumbUrl   string `json:"thumb_url,omitempty"`
    

    // moq 产品最小起订量
    
    Moq   int64 `json:"moq,omitempty"`
    

    // buy_now_url  产品下单链接
    
    BuyNowUrl   string `json:"buy_now_url,omitempty"`
    

    // currency 产品价格货币单位
    
    Currency   string `json:"currency,omitempty"`
    

}
*/

// GoodsSummary 
type GoodsSummary struct {

    // max_price_cent 产品最高价格
    MaxPriceCent   int64 `json:"max_price_cent,omitempty"`

    // id 产品id
    Id   string `json:"id,omitempty"`

    // unit 产品单位
    Unit   string `json:"unit,omitempty"`

    // detail_url 产品详情url
    DetailUrl   string `json:"detail_url,omitempty"`

    // supplier_id 产品对应供应商id
    SupplierId   string `json:"supplier_id,omitempty"`

    // subject 产品标题
    Subject   string `json:"subject,omitempty"`

    // min_price_cent 产品最低价
    MinPriceCent   int64 `json:"min_price_cent,omitempty"`

    // thumb_url 产品缩略图
    ThumbUrl   string `json:"thumb_url,omitempty"`

    // moq 产品最小起订量
    Moq   int64 `json:"moq,omitempty"`

    // buy_now_url  产品下单链接
    BuyNowUrl   string `json:"buy_now_url,omitempty"`

    // currency 产品价格货币单位
    Currency   string `json:"currency,omitempty"`

}

package scbp

// TopMatchedProductDto 
/* model for simplify = false
type TopMatchedProductDto struct {

    // 是否设置优先推广
    
    IsPreferential   string `json:"is_preferential,omitempty"`
    

    // 推广评分星级取值[0-5]
    
    QsStar   int64 `json:"qs_star,omitempty"`
    

    // 产品标题
    
    Subject   string `json:"subject,omitempty"`
    

    // 是否强制绑定
    
    IsForceMatch   string `json:"is_force_match,omitempty"`
    

    // 产品ID
    
    ProductId   int64 `json:"product_id,omitempty"`
    

    // Y:产品推广中,N:产品未推广
    
    IsOffer   string `json:"is_offer,omitempty"`
    

}
*/

// TopMatchedProductDto 
type TopMatchedProductDto struct {

    // 是否设置优先推广
    IsPreferential   string `json:"is_preferential,omitempty"`

    // 推广评分星级取值[0-5]
    QsStar   int64 `json:"qs_star,omitempty"`

    // 产品标题
    Subject   string `json:"subject,omitempty"`

    // 是否强制绑定
    IsForceMatch   string `json:"is_force_match,omitempty"`

    // 产品ID
    ProductId   int64 `json:"product_id,omitempty"`

    // Y:产品推广中,N:产品未推广
    IsOffer   string `json:"is_offer,omitempty"`

}

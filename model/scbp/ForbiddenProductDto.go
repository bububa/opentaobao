package scbp

// ForbiddenProductDto 
/* model for simplify = false
type ForbiddenProductDto struct {

    // 产品id
    
    ProductId   int64 `json:"product_id,omitempty"`
    

    // 状态
    
    Status   int64 `json:"status,omitempty"`
    

    // 标题
    
    Subject   string `json:"subject,omitempty"`
    

    // 图片地址
    
    ImgUrl   string `json:"img_url,omitempty"`
    

}
*/

// ForbiddenProductDto 
type ForbiddenProductDto struct {

    // 产品id
    ProductId   int64 `json:"product_id,omitempty"`

    // 状态
    Status   int64 `json:"status,omitempty"`

    // 标题
    Subject   string `json:"subject,omitempty"`

    // 图片地址
    ImgUrl   string `json:"img_url,omitempty"`

}

package scbp

// TopProductDto 
/* model for simplify = false
type TopProductDto struct {

    // 产品标题，最大长度256个字符
    
    Subject   string `json:"subject,omitempty"`
    

    // 产品推广状态，取值[disabled,enabled]
    
    Status   string `json:"status,omitempty"`
    

    // 产品ID
    
    ProductId   int64 `json:"product_id,omitempty"`
    

}
*/

// TopProductDto 
type TopProductDto struct {

    // 产品标题，最大长度256个字符
    Subject   string `json:"subject,omitempty"`

    // 产品推广状态，取值[disabled,enabled]
    Status   string `json:"status,omitempty"`

    // 产品ID
    ProductId   int64 `json:"product_id,omitempty"`

}

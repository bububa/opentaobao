package icbu

// ProductTopPublishRequest 
/* model for simplify = false
type ProductTopPublishRequest struct {

    // 类目id
    
    CatId   int64 `json:"cat_id,omitempty"`
    

    // 返回文案的语种，支持en_US,zh,zh_TW
    
    Language   string `json:"language,omitempty"`
    

    // 商品的具体数据信息
    
    Xml   string `json:"xml,omitempty"`
    

    // 商品明文id
    
    ProductId   int64 `json:"product_id,omitempty"`
    

}
*/

// ProductTopPublishRequest 
type ProductTopPublishRequest struct {

    // 类目id
    CatId   int64 `json:"cat_id,omitempty"`

    // 返回文案的语种，支持en_US,zh,zh_TW
    Language   string `json:"language,omitempty"`

    // 商品的具体数据信息
    Xml   string `json:"xml,omitempty"`

    // 商品明文id
    ProductId   int64 `json:"product_id,omitempty"`

}

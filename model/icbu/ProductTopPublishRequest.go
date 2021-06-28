package icbu

// ProductTopPublishRequest 
type ProductTopPublishRequest struct {

    // 类目id
    
    CatId   int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
    

    // 返回文案的语种，支持en_US,zh,zh_TW
    
    Language   string `json:"language,omitempty" xml:"language,omitempty"`
    

    // 商品的具体数据信息
    
    Xml   string `json:"xml,omitempty" xml:"xml,omitempty"`
    

    // 商品明文id
    
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
    

}

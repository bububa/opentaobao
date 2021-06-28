package scbp

// TopP4pQuickCampaignProductView 
/* model for simplify = false
type TopP4pQuickCampaignProductView struct {

    // 商品最近7天效果数据
    
    Effect7d  *struct {
        Effect7d  *Effect7d `json:"effect7d,omitempty"`
    } `json:"effect7d,omitempty"`
    

    // 商品名
    
    ProductName   string `json:"product_name,omitempty"`
    

    // 商品ID
    
    ProductId   int64 `json:"product_id,omitempty"`
    

    // 产品状态（0暂停，1推广中，-2商品下架）
    
    DisplayStatus   int64 `json:"display_status,omitempty"`
    

}
*/

// TopP4pQuickCampaignProductView 
type TopP4pQuickCampaignProductView struct {

    // 商品最近7天效果数据
    Effect7d   *Effect7d `json:"effect7d,omitempty"`

    // 商品名
    ProductName   string `json:"product_name,omitempty"`

    // 商品ID
    ProductId   int64 `json:"product_id,omitempty"`

    // 产品状态（0暂停，1推广中，-2商品下架）
    DisplayStatus   int64 `json:"display_status,omitempty"`

}

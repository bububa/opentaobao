package ascpchannel

// Products 
type Products struct {
    // 经营模式
    SalesModes   []string `json:"sales_modes,omitempty" xml:"sales_modes>string,omitempty"`
    // 图片
    Pictures   []string `json:"pictures,omitempty" xml:"pictures>string,omitempty"`
    // 产品标题
    ProductTitle   string `json:"product_title,omitempty" xml:"product_title,omitempty"`
    // 产品id
    ProductId   string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

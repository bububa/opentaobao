package ascpchannel

// ProductDetailDTO 
type ProductDetailDTO struct {

    // sku 列表
    
    ProductSkuDetail   []ProductSkuDetailDTO `json:"product_sku_detail,omitempty" xml:"product_sku_detail,omitempty"`
    

    // 透明白底图
    
    WhiteBgPicture   string `json:"white_bg_picture,omitempty" xml:"white_bg_picture,omitempty"`
    

    // 产品 id
    
    ProductId   string `json:"product_id,omitempty" xml:"product_id,omitempty"`
    

    // 富文本详描
    
    DescRichText   string `json:"desc_rich_text,omitempty" xml:"desc_rich_text,omitempty"`
    

    // 经营模式
    
    SalesModes   []string `json:"sales_modes,omitempty" xml:"sales_modes>string,omitempty"`
    

    // 二级渠道编码
    
    SubChannelCode   string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
    

    // 产品图片链接
    
    Pictures   []string `json:"pictures,omitempty" xml:"pictures>string,omitempty"`
    

    // 产品标题
    
    ProductTitle   string `json:"product_title,omitempty" xml:"product_title,omitempty"`
    

    // 产品销售属性
    
    Properties   []string `json:"properties,omitempty" xml:"properties>string,omitempty"`
    

    // 类目
    
    Category   string `json:"category,omitempty" xml:"category,omitempty"`
    

    // 品牌
    
    Brand   string `json:"brand,omitempty" xml:"brand,omitempty"`
    

    // 渠道编码
    
    ChannelCode   string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
    

    // 品牌 id
    
    BrandId   string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    

    // 类目 id
    
    CategoryId   string `json:"category_id,omitempty" xml:"category_id,omitempty"`
    

}

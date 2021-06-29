package tmallchannel

// ProductTopDto 
type ProductTopDto struct {

    // 产品线ID
    
    ProductLineId   int64 `json:"product_line_id,omitempty" xml:"product_line_id,omitempty"`
    

    // 产品编码
    
    ProductNumber   string `json:"product_number,omitempty" xml:"product_number,omitempty"`
    

    // 产品描述地址
    
    DescPath   string `json:"desc_path,omitempty" xml:"desc_path,omitempty"`
    

    // 基准价
    
    StandardPrice   int64 `json:"standard_price,omitempty" xml:"standard_price,omitempty"`
    

    // 标题
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 产品Id
    
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
    

    // 没有sku的情况下，产品对应的后端商品id
    
    ScItemId   int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
    

    // spuId
    
    SpuId   int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
    

    // 供应商Id
    
    SupplierId   int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
    

    // 类目Id
    
    CategoryId   int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
    

    // sku列表
    
    SkuList   []ProductSkuTopDto `json:"sku_list,omitempty" xml:"sku_list,omitempty"`
    

}

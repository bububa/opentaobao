package openmall

// TopSaleAreaVo 
type TopSaleAreaVo struct {

    // 商品SKU ID
    
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    

    // 可售区域信息，JSON数组格式的字符串。区域的层次用树状结构表示。 包含的字段有：     areaId：区域码; 	subSaleArea：所属子区域的可售信息; 	 树的叶子节点表示可售区域。 树的枝节点只是聚合可售区域，表达区域的层次关系。
    
    SaleAreaInfo   string `json:"sale_area_info,omitempty" xml:"sale_area_info,omitempty"`
    

}

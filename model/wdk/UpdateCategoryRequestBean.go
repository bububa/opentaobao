package wdk

// UpdateCategoryRequestBean 
type UpdateCategoryRequestBean struct {

    // 机构编码
    
    OrgCode   string `json:"org_code,omitempty" xml:"org_code,omitempty"`
    

    // 商品编码
    
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    

    // forest类目ID
    
    ForestId   string `json:"forest_id,omitempty" xml:"forest_id,omitempty"`
    

    // 盒马 后台类目
    
    BackCatCode   string `json:"back_cat_code,omitempty" xml:"back_cat_code,omitempty"`
    

}

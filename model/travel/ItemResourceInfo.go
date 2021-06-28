package travel

// ItemResourceInfo 
/* model for simplify = false
type ItemResourceInfo struct {

    // 对应的说明
    
    Desc   string `json:"desc,omitempty"`
    

    // 1-保险 2-餐饮 3-租车 4-签证 5-购物点，预留  6-赠品，预留 7-券，预留  8-其他费用
    
    Type   int64 `json:"type,omitempty"`
    

    // 关联的套餐id
    
    RelatedPackageId   int64 `json:"related_package_id,omitempty"`
    

}
*/

// ItemResourceInfo 
type ItemResourceInfo struct {

    // 对应的说明
    Desc   string `json:"desc,omitempty"`

    // 1-保险 2-餐饮 3-租车 4-签证 5-购物点，预留  6-赠品，预留 7-券，预留  8-其他费用
    Type   int64 `json:"type,omitempty"`

    // 关联的套餐id
    RelatedPackageId   int64 `json:"related_package_id,omitempty"`

}

package travel

// PontusTravelItemResourceInfo 
/* model for simplify = false
type PontusTravelItemResourceInfo struct {

    // 小于1500字符
    
    Desc   string `json:"desc,omitempty"`
    

    // 1-保险 2-餐饮 3-租车 4-签证 5-购物点，预留  6-赠品，预留 7-券，预留  8-其他费用
    
    Type   int64 `json:"type,omitempty"`
    

}
*/

// PontusTravelItemResourceInfo 
type PontusTravelItemResourceInfo struct {

    // 小于1500字符
    Desc   string `json:"desc,omitempty"`

    // 1-保险 2-餐饮 3-租车 4-签证 5-购物点，预留  6-赠品，预留 7-券，预留  8-其他费用
    Type   int64 `json:"type,omitempty"`

}

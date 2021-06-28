package waybill

// PackageItem 
/* model for simplify = false
type PackageItem struct {

    // 商品名称
    
    ItemName   string `json:"item_name,omitempty"`
    

    // 商品数量
    
    Count   int64 `json:"count,omitempty"`
    

}
*/

// PackageItem 
type PackageItem struct {

    // 商品名称
    ItemName   string `json:"item_name,omitempty"`

    // 商品数量
    Count   int64 `json:"count,omitempty"`

}

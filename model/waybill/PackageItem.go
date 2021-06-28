package waybill

// PackageItem 
type PackageItem struct {

    // 商品名称
    
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    

    // 商品数量
    
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    

}

package mos

// StoreInfo 
/* model for simplify = false
type StoreInfo struct {

    // 图片
    
    StorePic   string `json:"store_pic,omitempty"`
    

    // 专柜ID
    
    StoreId   int64 `json:"store_id,omitempty"`
    

    // 专柜名字
    
    StoreName   string `json:"store_name,omitempty"`
    

}
*/

// StoreInfo 
type StoreInfo struct {

    // 图片
    StorePic   string `json:"store_pic,omitempty"`

    // 专柜ID
    StoreId   int64 `json:"store_id,omitempty"`

    // 专柜名字
    StoreName   string `json:"store_name,omitempty"`

}

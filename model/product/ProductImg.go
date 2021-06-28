package product

// ProductImg 
/* model for simplify = false
type ProductImg struct {

    // 产品图片ID
    
    Id   int64 `json:"id,omitempty"`
    

    // 图片所属产品的ID
    
    ProductId   int64 `json:"product_id,omitempty"`
    

    // 图片地址.(绝对地址,格式:http://host/image_path)
    
    Url   string `json:"url,omitempty"`
    

    // 图片序号。产品里的图片展示顺序，数据越小越靠前。要求是正整数。
    
    Position   int64 `json:"position,omitempty"`
    

    // 添加时间.格式:yyyy-mm-dd hh:mm:ss
    
    Created   string `json:"created,omitempty"`
    

    // 修改时间.格式:yyyy-mm-dd hh:mm:ss
    
    Modified   string `json:"modified,omitempty"`
    

}
*/

// ProductImg 
type ProductImg struct {

    // 产品图片ID
    Id   int64 `json:"id,omitempty"`

    // 图片所属产品的ID
    ProductId   int64 `json:"product_id,omitempty"`

    // 图片地址.(绝对地址,格式:http://host/image_path)
    Url   string `json:"url,omitempty"`

    // 图片序号。产品里的图片展示顺序，数据越小越靠前。要求是正整数。
    Position   int64 `json:"position,omitempty"`

    // 添加时间.格式:yyyy-mm-dd hh:mm:ss
    Created   string `json:"created,omitempty"`

    // 修改时间.格式:yyyy-mm-dd hh:mm:ss
    Modified   string `json:"modified,omitempty"`

}

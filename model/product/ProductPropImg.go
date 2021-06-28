package product

// ProductPropImg 
type ProductPropImg struct {

    // 产品属性图片ID
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 图片所属产品的ID
    
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
    

    // 属性串(pid:vid),目前只有颜色属性.如:颜色:红色表示为　1627207:28326
    
    Props   string `json:"props,omitempty" xml:"props,omitempty"`
    

    // 图片地址.(绝对地址,格式:http://host/image_path)
    
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
    

    // 图片序号。产品里的图片展示顺序，数据越小越靠前。要求是正整数。
    
    Position   int64 `json:"position,omitempty" xml:"position,omitempty"`
    

    // 添加时间.格式:yyyy-mm-dd hh:mm:ss
    
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
    

    // 修改时间.格式:yyyy-mm-dd hh:mm:ss
    
    Modified   string `json:"modified,omitempty" xml:"modified,omitempty"`
    

}

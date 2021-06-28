package travel

// ProductHighLights 
/* model for simplify = false
type ProductHighLights struct {

    // 产品亮点标题
    
    Title   string `json:"title,omitempty"`
    

    // 产品亮点描述
    
    Desc   string `json:"desc,omitempty"`
    

    // 产品亮点图片
    
    PicUrls  struct {
        String  []string `json:"string,omitempty"`
    } `json:"pic_urls,omitempty"`
    

}
*/

// ProductHighLights 
type ProductHighLights struct {

    // 产品亮点标题
    Title   string `json:"title,omitempty"`

    // 产品亮点描述
    Desc   string `json:"desc,omitempty"`

    // 产品亮点图片
    PicUrls   []string `json:"pic_urls,omitempty"`

}

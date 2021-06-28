package travel

// HighLights 
/* model for simplify = false
type HighLights struct {

    // 亮点标题
    
    Title   string `json:"title,omitempty"`
    

    // 亮点描述
    
    Desc   string `json:"desc,omitempty"`
    

    // 图片列表
    
    PicUrls  struct {
        String  []string `json:"string,omitempty"`
    } `json:"pic_urls,omitempty"`
    

}
*/

// HighLights 
type HighLights struct {

    // 亮点标题
    Title   string `json:"title,omitempty"`

    // 亮点描述
    Desc   string `json:"desc,omitempty"`

    // 图片列表
    PicUrls   []string `json:"pic_urls,omitempty"`

}

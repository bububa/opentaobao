package baichuan

// BcTaoPasswordResult 
/* model for simplify = false
type BcTaoPasswordResult struct {

    // 淘口令-文案
    
    Content   string `json:"content,omitempty"`
    

    // 淘口令-宝贝
    
    Title   string `json:"title,omitempty"`
    

    // 图片url
    
    PicUrl   string `json:"pic_url,omitempty"`
    

    // 查询结果code
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // 查询结果code
    
    ResultMessage   string `json:"result_message,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 跳转url(长链)
    
    Url   string `json:"url,omitempty"`
    

    // 如果是宝贝，则为宝贝价格
    
    Price   string `json:"price,omitempty"`
    

    // nativeUrl
    
    NativeUrl   string `json:"native_url,omitempty"`
    

    // thumbPicUrl
    
    ThumbPicUrl   string `json:"thumb_pic_url,omitempty"`
    

}
*/

// BcTaoPasswordResult 
type BcTaoPasswordResult struct {

    // 淘口令-文案
    Content   string `json:"content,omitempty"`

    // 淘口令-宝贝
    Title   string `json:"title,omitempty"`

    // 图片url
    PicUrl   string `json:"pic_url,omitempty"`

    // 查询结果code
    ResultCode   string `json:"result_code,omitempty"`

    // 查询结果code
    ResultMessage   string `json:"result_message,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 跳转url(长链)
    Url   string `json:"url,omitempty"`

    // 如果是宝贝，则为宝贝价格
    Price   string `json:"price,omitempty"`

    // nativeUrl
    NativeUrl   string `json:"native_url,omitempty"`

    // thumbPicUrl
    ThumbPicUrl   string `json:"thumb_pic_url,omitempty"`

}

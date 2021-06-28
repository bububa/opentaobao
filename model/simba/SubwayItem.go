package simba

// SubwayItem 
/* model for simplify = false
type SubwayItem struct {

    // 商品信息在外部系统(淘宝主站)的数字id
    
    NumId   int64 `json:"num_id,omitempty"`
    

    // 商品信息在外部系统（淘宝主站）的价格
    
    Price   string `json:"price,omitempty"`
    

    // 商品信息在外部系统（淘宝主站）的标题
    
    Title   string `json:"title,omitempty"`
    

    // 扩展属性对象
    
    ExtraAttributes  *struct {
        ExtraAttributes  *ExtraAttributes `json:"extra_attributes,omitempty"`
    } `json:"extra_attributes,omitempty"`
    

    // imgUrl
    
    ImgUrl   string `json:"img_url,omitempty"`
    

}
*/

// SubwayItem 
type SubwayItem struct {

    // 商品信息在外部系统(淘宝主站)的数字id
    NumId   int64 `json:"num_id,omitempty"`

    // 商品信息在外部系统（淘宝主站）的价格
    Price   string `json:"price,omitempty"`

    // 商品信息在外部系统（淘宝主站）的标题
    Title   string `json:"title,omitempty"`

    // 扩展属性对象
    ExtraAttributes   *ExtraAttributes `json:"extra_attributes,omitempty"`

    // imgUrl
    ImgUrl   string `json:"img_url,omitempty"`

}

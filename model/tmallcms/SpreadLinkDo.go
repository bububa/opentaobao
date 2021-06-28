package tmallcms

// SpreadLinkDo 
/* model for simplify = false
type SpreadLinkDo struct {

    // ID
    
    Id   int64 `json:"id,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 推广文案
    
    Adword   string `json:"adword,omitempty"`
    

    // 推广信息
    
    Adinfo   string `json:"adinfo,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 状态
    
    Status   int64 `json:"status,omitempty"`
    

    // 商家ID
    
    SellerId   int64 `json:"seller_id,omitempty"`
    

    // 微信推广链接
    
    Adwxurl   string `json:"adwxurl,omitempty"`
    

    // 二维码地址
    
    Qrcode   string `json:"qrcode,omitempty"`
    

    // 推广链接
    
    Adurl   string `json:"adurl,omitempty"`
    

    // 类型
    
    Type   int64 `json:"type,omitempty"`
    

    // 源链接
    
    Url   string `json:"url,omitempty"`
    

}
*/

// SpreadLinkDo 
type SpreadLinkDo struct {

    // ID
    Id   int64 `json:"id,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 推广文案
    Adword   string `json:"adword,omitempty"`

    // 推广信息
    Adinfo   string `json:"adinfo,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 状态
    Status   int64 `json:"status,omitempty"`

    // 商家ID
    SellerId   int64 `json:"seller_id,omitempty"`

    // 微信推广链接
    Adwxurl   string `json:"adwxurl,omitempty"`

    // 二维码地址
    Qrcode   string `json:"qrcode,omitempty"`

    // 推广链接
    Adurl   string `json:"adurl,omitempty"`

    // 类型
    Type   int64 `json:"type,omitempty"`

    // 源链接
    Url   string `json:"url,omitempty"`

}

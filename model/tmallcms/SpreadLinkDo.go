package tmallcms

// SpreadLinkDo 
type SpreadLinkDo struct {
    // ID
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 推广文案
    Adword   string `json:"adword,omitempty" xml:"adword,omitempty"`
    // 推广信息
    Adinfo   string `json:"adinfo,omitempty" xml:"adinfo,omitempty"`
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 状态
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 商家ID
    SellerId   int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
    // 微信推广链接
    Adwxurl   string `json:"adwxurl,omitempty" xml:"adwxurl,omitempty"`
    // 二维码地址
    Qrcode   string `json:"qrcode,omitempty" xml:"qrcode,omitempty"`
    // 推广链接
    Adurl   string `json:"adurl,omitempty" xml:"adurl,omitempty"`
    // 类型
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    // 源链接
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
}

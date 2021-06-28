package jst

// OfficialAccountInfoReportRequest 
/* model for simplify = false
type OfficialAccountInfoReportRequest struct {

    // ISV给商家分配的通道号码后缀，可传多个，逗号分隔
    
    Suffixes   string `json:"suffixes,omitempty"`
    

    // 白底logo图片的资源Url地址，正方形，400x400，png格式
    
    ShopLogo   string `json:"shop_logo,omitempty"`
    

    // 店铺名
    
    ShopName   string `json:"shop_name,omitempty"`
    

    // ADD：新增，MODIFY：修改
    
    Operation   string `json:"operation,omitempty"`
    

    // 短信签名
    
    SmsSign   string `json:"sms_sign,omitempty"`
    

}
*/

// OfficialAccountInfoReportRequest 
type OfficialAccountInfoReportRequest struct {

    // ISV给商家分配的通道号码后缀，可传多个，逗号分隔
    Suffixes   string `json:"suffixes,omitempty"`

    // 白底logo图片的资源Url地址，正方形，400x400，png格式
    ShopLogo   string `json:"shop_logo,omitempty"`

    // 店铺名
    ShopName   string `json:"shop_name,omitempty"`

    // ADD：新增，MODIFY：修改
    Operation   string `json:"operation,omitempty"`

    // 短信签名
    SmsSign   string `json:"sms_sign,omitempty"`

}

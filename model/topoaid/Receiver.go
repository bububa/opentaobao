package topoaid

// Receiver 
type Receiver struct {

    // 收件人的姓名
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 收件人的手机号
    
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    

    // 收件人的详细地址
    
    AddressDetail   string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
    

    // 收件人的电话号码
    
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    

    // 收件人ID (Open Addressee ID)，长度在128个字符之内。
    
    Oaid   string `json:"oaid,omitempty" xml:"oaid,omitempty"`
    

    // 交易编号
    
    Tid   string `json:"tid,omitempty" xml:"tid,omitempty"`
    

    // 收货人街道地址
    
    Town   string `json:"town,omitempty" xml:"town,omitempty"`
    

    // 收货人的所在地区
    
    District   string `json:"district,omitempty" xml:"district,omitempty"`
    

    // 收货人的所在城市
    
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    

    // 收货人的所在省份
    
    State   string `json:"state,omitempty" xml:"state,omitempty"`
    

    // 收货人国籍
    
    Country   string `json:"country,omitempty" xml:"country,omitempty"`
    

    // oaid是否和tid当前的oaid匹配。true：匹配，false：不匹配。当不匹配时，建议通过taobao.trade.fullinfo.get获取最新的oaid。
    
    Matched   bool `json:"matched,omitempty" xml:"matched,omitempty"`
    

}

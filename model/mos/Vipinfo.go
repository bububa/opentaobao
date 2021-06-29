package mos

// Vipinfo 
type Vipinfo struct {
    // 会员账号
    Integral   string `json:"integral,omitempty" xml:"integral,omitempty"`
    // 会员卡号
    Cardid   string `json:"cardid,omitempty" xml:"cardid,omitempty"`
    // 会员卡类别
    Hyklb   string `json:"hyklb,omitempty" xml:"hyklb,omitempty"`
    // 折扣卡类别
    Zkklb   string `json:"zkklb,omitempty" xml:"zkklb,omitempty"`
    // 会员手机号
    Memo   string `json:"memo,omitempty" xml:"memo,omitempty"`
    // 喵街会员类别
    Alicardtype   string `json:"alicardtype,omitempty" xml:"alicardtype,omitempty"`
}

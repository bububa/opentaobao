package mos

// SumStr 
type SumStr struct {
    // 小票交易类型
    N0609   string `json:"n0609,omitempty" xml:"n0609,omitempty"`
    // 小票交易总金额
    N0606   string `json:"n0606,omitempty" xml:"n0606,omitempty"`
    // 小票总折扣
    N0607   string `json:"n0607,omitempty" xml:"n0607,omitempty"`
    // 预购小票标志
    N0610   string `json:"n0610,omitempty" xml:"n0610,omitempty"`
    // 会员卡号
    N0612   string `json:"n0612,omitempty" xml:"n0612,omitempty"`
    // 会员卡类别
    N0613   string `json:"n0613,omitempty" xml:"n0613,omitempty"`
    // 小票号
    Seqno   string `json:"seqno,omitempty" xml:"seqno,omitempty"`
    // 门店号
    Storeno   string `json:"storeno,omitempty" xml:"storeno,omitempty"`
}

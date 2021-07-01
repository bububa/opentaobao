package mos

// Payment 结构体
type Payment struct {
	// 支付方式大类
	N0703 string `json:"n0703,omitempty" xml:"n0703,omitempty"`
	// 支付方式金额
	N0704 string `json:"n0704,omitempty" xml:"n0704,omitempty"`
	// 券号，订单号等
	N0705 string `json:"n0705,omitempty" xml:"n0705,omitempty"`
	// 支付方式小类
	N0710 string `json:"n0710,omitempty" xml:"n0710,omitempty"`
	// 支付方式行号
	N0711 string `json:"n0711,omitempty" xml:"n0711,omitempty"`
	// 支付方式名称
	Fkfsname string `json:"fkfsname,omitempty" xml:"fkfsname,omitempty"`
}

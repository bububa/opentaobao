package ascp

// ConfirmPackages 结构体
type ConfirmPackages struct {
	// 运单内所包含的所有货品（与翱象对接的货品编码）
	ScItems []ScItems `json:"sc_items,omitempty" xml:"sc_items>sc_items,omitempty"`
	// 物流公司编码(SF=顺丰、EMS=标准快递、EYB=经济快件、ZJS=宅急送、YTO=圆通、ZTO=中通 (ZTO)、HTKY=百世汇通、 UC=优速、STO=申通、TTKDEX=天天快递、QFKD=全峰、FAST=快捷、POSTB=邮政小包、GTO=国通、YUNDA=韵达、JD=京东配送、DD=当当宅配、 AMAZON=亚马逊物流、OTHERS=其他；只传英文编码)
	LogisticsCode string `json:"logistics_code,omitempty" xml:"logistics_code,omitempty"`
	// 运单号
	ExpressCode string `json:"express_code,omitempty" xml:"express_code,omitempty"`
}

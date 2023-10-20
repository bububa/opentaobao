package ascpchannel

// AlibabaAscpIndustryWaybillPreAcceptData 结构体
type AlibabaAscpIndustryWaybillPreAcceptData struct {
	// 运单号
	ExpressNo string `json:"express_no,omitempty" xml:"express_no,omitempty"`
	// 物流品牌
	LogisticsBrand string `json:"logistics_brand,omitempty" xml:"logistics_brand,omitempty"`
}

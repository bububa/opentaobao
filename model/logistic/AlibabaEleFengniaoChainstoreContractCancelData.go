package logistic

// AlibabaEleFengniaoChainstoreContractCancelData 结构体
type AlibabaEleFengniaoChainstoreContractCancelData struct {
	// 门店code
	ChainstoreCodes []string `json:"chainstore_codes,omitempty" xml:"chainstore_codes>string,omitempty"`
	// appid
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 商户code
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
}

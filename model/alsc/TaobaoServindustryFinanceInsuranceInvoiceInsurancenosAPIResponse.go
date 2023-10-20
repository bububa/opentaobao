package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoservindustryfinanceinsuranceinvoiceinsurancenosAPIResponse 商家查询本次开票的保险单号 API返回值
// taobao.servindustry.finance.insurance.invoice.insurancenos
//
// 商家查询本次开票的保险单号
type TaobaoservindustryfinanceinsuranceinvoiceinsurancenosAPIResponse struct {
	model.CommonResponse
	TaobaoservindustryfinanceinsuranceinvoiceinsurancenosAPIResponseModel
}

// TaobaoservindustryfinanceinsuranceinvoiceinsurancenosAPIResponseModel is 商家查询本次开票的保险单号 成功返回结果
type TaobaoservindustryfinanceinsuranceinvoiceinsurancenosAPIResponseModel struct {
	XMLName xml.Name `xml:"servindustry_finance_insurance_invoice_insurancenos_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *TaobaoservindustryfinanceinsuranceinvoiceinsurancenosResult `json:"result,omitempty" xml:"result,omitempty"`
}

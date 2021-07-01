package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceUnitorderCheckAPIResponse
服务商订购单上传核对 API返回值
alibaba.einvoice.unitorder.check

开票服务商回传收到的订购单用于电子发票平台核对 */
type AlibabaEinvoiceUnitorderCheckAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceUnitorderCheckAPIResponseModel
}

// AlibabaEinvoiceUnitorderCheckAPIResponseModel is 服务商订购单上传核对 成功返回结果
type AlibabaEinvoiceUnitorderCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_unitorder_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 上传结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

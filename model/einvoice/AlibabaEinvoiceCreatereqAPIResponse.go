package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceCreatereqAPIResponse
ERP开票请求接口 API返回值
alibaba.einvoice.createreq

ERP发起开票请求 */
type AlibabaEinvoiceCreatereqAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceCreatereqAPIResponseModel
}

// AlibabaEinvoiceCreatereqAPIResponseModel is ERP开票请求接口 成功返回结果
type AlibabaEinvoiceCreatereqAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_createreq_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 开票信息是否成功接受
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

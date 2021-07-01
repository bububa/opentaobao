package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceClosereqAPIResponse
关闭开票失败请求（失败列表可重试） API返回值
alibaba.einvoice.closereq

关闭失败开票请求，避免造成重复开票 */
type AlibabaEinvoiceClosereqAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceClosereqAPIResponseModel
}

// AlibabaEinvoiceClosereqAPIResponseModel is 关闭开票失败请求（失败列表可重试） 成功返回结果
type AlibabaEinvoiceClosereqAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_closereq_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关闭是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoicePaperReturnAPIResponse
纸质发票结果回传 API返回值
alibaba.einvoice.paper.return

纸质发票结果回传 */
type AlibabaEinvoicePaperReturnAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoicePaperReturnAPIResponseModel
}

// AlibabaEinvoicePaperReturnAPIResponseModel is 纸质发票结果回传 成功返回结果
type AlibabaEinvoicePaperReturnAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_paper_return_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务端接收开票回传数据的结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

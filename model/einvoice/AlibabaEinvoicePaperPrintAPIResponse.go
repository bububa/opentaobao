package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicepaperprintAPIResponse 纸票打印接口 API返回值
// alibaba.einvoice.paper.print
//
// 打印一张已开具成功的纸票
type AlibabaeinvoicepaperprintAPIResponse struct {
	model.CommonResponse
	AlibabaeinvoicepaperprintAPIResponseModel
}

// AlibabaeinvoicepaperprintAPIResponseModel is 纸票打印接口 成功返回结果
type AlibabaeinvoicepaperprintAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_paper_print_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果，打印结果tmc消息alibaba_invoice_PaperOpsReturn异步通知
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLocalLogisticLabelPrintAPIResponse 物流打印面单 API返回值
// aliexpress.local.logistic.label.print
//
// 物流打印面单
type AliexpressLocalLogisticLabelPrintAPIResponse struct {
	model.CommonResponse
	AliexpressLocalLogisticLabelPrintAPIResponseModel
}

// AliexpressLocalLogisticLabelPrintAPIResponseModel is 物流打印面单 成功返回结果
type AliexpressLocalLogisticLabelPrintAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_local_logistic_label_print_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorResultMessage string `json:"error_result_message,omitempty" xml:"error_result_message,omitempty"`
	// 错误码
	ErrorResultCode string `json:"error_result_code,omitempty" xml:"error_result_code,omitempty"`
	// 出参对象
	Data *LabelDto `json:"data,omitempty" xml:"data,omitempty"`
	// 接口调用状态
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

package logistic

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AliexpressLocalLogisticLabelPrintAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressLocalLogisticLabelPrintAPIResponseModel).Reset()
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

// Reset 清空结构体
func (m *AliexpressLocalLogisticLabelPrintAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorResultMessage = ""
	m.ErrorResultCode = ""
	m.Data = nil
	m.IsSuccess = false
}

var poolAliexpressLocalLogisticLabelPrintAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressLocalLogisticLabelPrintAPIResponse)
	},
}

// GetAliexpressLocalLogisticLabelPrintAPIResponse 从 sync.Pool 获取 AliexpressLocalLogisticLabelPrintAPIResponse
func GetAliexpressLocalLogisticLabelPrintAPIResponse() *AliexpressLocalLogisticLabelPrintAPIResponse {
	return poolAliexpressLocalLogisticLabelPrintAPIResponse.Get().(*AliexpressLocalLogisticLabelPrintAPIResponse)
}

// ReleaseAliexpressLocalLogisticLabelPrintAPIResponse 将 AliexpressLocalLogisticLabelPrintAPIResponse 保存到 sync.Pool
func ReleaseAliexpressLocalLogisticLabelPrintAPIResponse(v *AliexpressLocalLogisticLabelPrintAPIResponse) {
	v.Reset()
	poolAliexpressLocalLogisticLabelPrintAPIResponse.Put(v)
}

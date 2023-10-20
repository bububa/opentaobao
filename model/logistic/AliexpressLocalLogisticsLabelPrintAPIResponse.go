package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLocalLogisticsLabelPrintAPIResponse print label API返回值
// aliexpress.local.logistics.label.print
//
// print label
type AliexpressLocalLogisticsLabelPrintAPIResponse struct {
	model.CommonResponse
	AliexpressLocalLogisticsLabelPrintAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressLocalLogisticsLabelPrintAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressLocalLogisticsLabelPrintAPIResponseModel).Reset()
}

// AliexpressLocalLogisticsLabelPrintAPIResponseModel is print label 成功返回结果
type AliexpressLocalLogisticsLabelPrintAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_local_logistics_label_print_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// error message
	ErrorResultMessage string `json:"error_result_message,omitempty" xml:"error_result_message,omitempty"`
	// error code
	ErrorResultCode string `json:"error_result_code,omitempty" xml:"error_result_code,omitempty"`
	// data
	Data *LabelDto `json:"data,omitempty" xml:"data,omitempty"`
	// is success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressLocalLogisticsLabelPrintAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorResultMessage = ""
	m.ErrorResultCode = ""
	m.Data = nil
	m.IsSuccess = false
}

var poolAliexpressLocalLogisticsLabelPrintAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressLocalLogisticsLabelPrintAPIResponse)
	},
}

// GetAliexpressLocalLogisticsLabelPrintAPIResponse 从 sync.Pool 获取 AliexpressLocalLogisticsLabelPrintAPIResponse
func GetAliexpressLocalLogisticsLabelPrintAPIResponse() *AliexpressLocalLogisticsLabelPrintAPIResponse {
	return poolAliexpressLocalLogisticsLabelPrintAPIResponse.Get().(*AliexpressLocalLogisticsLabelPrintAPIResponse)
}

// ReleaseAliexpressLocalLogisticsLabelPrintAPIResponse 将 AliexpressLocalLogisticsLabelPrintAPIResponse 保存到 sync.Pool
func ReleaseAliexpressLocalLogisticsLabelPrintAPIResponse(v *AliexpressLocalLogisticsLabelPrintAPIResponse) {
	v.Reset()
	poolAliexpressLocalLogisticsLabelPrintAPIResponse.Put(v)
}

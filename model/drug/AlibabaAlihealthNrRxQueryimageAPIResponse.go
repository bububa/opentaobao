package drug

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthNrRxQueryimageAPIResponse o2o查看处方图片 API返回值
// alibaba.alihealth.nr.rx.queryimage
//
// o2o商家查看处方图片，包括电子图片与纸质图片
type AlibabaAlihealthNrRxQueryimageAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthNrRxQueryimageAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthNrRxQueryimageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthNrRxQueryimageAPIResponseModel).Reset()
}

// AlibabaAlihealthNrRxQueryimageAPIResponseModel is o2o查看处方图片 成功返回结果
type AlibabaAlihealthNrRxQueryimageAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_nr_rx_queryimage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误编码
	ErrorInfoCode string `json:"error_info_code,omitempty" xml:"error_info_code,omitempty"`
	// 错误描述
	ErrorInfoMsg string `json:"error_info_msg,omitempty" xml:"error_info_msg,omitempty"`
	// 处方图片url，多张图片用逗号分隔
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 成功或失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthNrRxQueryimageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorInfoCode = ""
	m.ErrorInfoMsg = ""
	m.Result = ""
	m.IsSuccess = false
}

var poolAlibabaAlihealthNrRxQueryimageAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthNrRxQueryimageAPIResponse)
	},
}

// GetAlibabaAlihealthNrRxQueryimageAPIResponse 从 sync.Pool 获取 AlibabaAlihealthNrRxQueryimageAPIResponse
func GetAlibabaAlihealthNrRxQueryimageAPIResponse() *AlibabaAlihealthNrRxQueryimageAPIResponse {
	return poolAlibabaAlihealthNrRxQueryimageAPIResponse.Get().(*AlibabaAlihealthNrRxQueryimageAPIResponse)
}

// ReleaseAlibabaAlihealthNrRxQueryimageAPIResponse 将 AlibabaAlihealthNrRxQueryimageAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthNrRxQueryimageAPIResponse(v *AlibabaAlihealthNrRxQueryimageAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthNrRxQueryimageAPIResponse.Put(v)
}

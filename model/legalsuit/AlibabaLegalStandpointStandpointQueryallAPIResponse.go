package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointStandpointQueryallAPIResponse 滑动查询口径 API返回值
// alibaba.legal.standpoint.standpoint.queryall
//
// 滑动查询口径
type AlibabaLegalStandpointStandpointQueryallAPIResponse struct {
	model.CommonResponse
	AlibabaLegalStandpointStandpointQueryallAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointStandpointQueryallAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalStandpointStandpointQueryallAPIResponseModel).Reset()
}

// AlibabaLegalStandpointStandpointQueryallAPIResponseModel is 滑动查询口径 成功返回结果
type AlibabaLegalStandpointStandpointQueryallAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_standpoint_standpoint_queryall_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 口径
	Content []StandpointOutPutDto `json:"content,omitempty" xml:"content>standpoint_out_put_dto,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误编码
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 是否成功
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointStandpointQueryallAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Content = m.Content[:0]
	m.ErrorMsg = ""
	m.ErrorCodeRes = 0
	m.SuccessRes = false
}

var poolAlibabaLegalStandpointStandpointQueryallAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalStandpointStandpointQueryallAPIResponse)
	},
}

// GetAlibabaLegalStandpointStandpointQueryallAPIResponse 从 sync.Pool 获取 AlibabaLegalStandpointStandpointQueryallAPIResponse
func GetAlibabaLegalStandpointStandpointQueryallAPIResponse() *AlibabaLegalStandpointStandpointQueryallAPIResponse {
	return poolAlibabaLegalStandpointStandpointQueryallAPIResponse.Get().(*AlibabaLegalStandpointStandpointQueryallAPIResponse)
}

// ReleaseAlibabaLegalStandpointStandpointQueryallAPIResponse 将 AlibabaLegalStandpointStandpointQueryallAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalStandpointStandpointQueryallAPIResponse(v *AlibabaLegalStandpointStandpointQueryallAPIResponse) {
	v.Reset()
	poolAlibabaLegalStandpointStandpointQueryallAPIResponse.Put(v)
}

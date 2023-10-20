package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTmsCutConfirmAPIResponse 配拦截失败CP确认结果并回告 API返回值
// alibaba.alihealth.tms.cut.confirm
//
// 配拦截失败CP确认结果并回告
type AlibabaAlihealthTmsCutConfirmAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthTmsCutConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthTmsCutConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthTmsCutConfirmAPIResponseModel).Reset()
}

// AlibabaAlihealthTmsCutConfirmAPIResponseModel is 配拦截失败CP确认结果并回告 成功返回结果
type AlibabaAlihealthTmsCutConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tms_cut_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAlihealthTmsCutConfirmResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthTmsCutConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthTmsCutConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTmsCutConfirmAPIResponse)
	},
}

// GetAlibabaAlihealthTmsCutConfirmAPIResponse 从 sync.Pool 获取 AlibabaAlihealthTmsCutConfirmAPIResponse
func GetAlibabaAlihealthTmsCutConfirmAPIResponse() *AlibabaAlihealthTmsCutConfirmAPIResponse {
	return poolAlibabaAlihealthTmsCutConfirmAPIResponse.Get().(*AlibabaAlihealthTmsCutConfirmAPIResponse)
}

// ReleaseAlibabaAlihealthTmsCutConfirmAPIResponse 将 AlibabaAlihealthTmsCutConfirmAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthTmsCutConfirmAPIResponse(v *AlibabaAlihealthTmsCutConfirmAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthTmsCutConfirmAPIResponse.Put(v)
}

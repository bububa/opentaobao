package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterWorkcardCancelAPIResponse 服务平台工单取消接口 API返回值
// alibaba.servicecenter.workcard.cancel
//
// 取消服务工单
type AlibabaServicecenterWorkcardCancelAPIResponse struct {
	model.CommonResponse
	AlibabaServicecenterWorkcardCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaServicecenterWorkcardCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaServicecenterWorkcardCancelAPIResponseModel).Reset()
}

// AlibabaServicecenterWorkcardCancelAPIResponseModel is 服务平台工单取消接口 成功返回结果
type AlibabaServicecenterWorkcardCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_workcard_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaServicecenterWorkcardCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaServicecenterWorkcardCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterWorkcardCancelAPIResponse)
	},
}

// GetAlibabaServicecenterWorkcardCancelAPIResponse 从 sync.Pool 获取 AlibabaServicecenterWorkcardCancelAPIResponse
func GetAlibabaServicecenterWorkcardCancelAPIResponse() *AlibabaServicecenterWorkcardCancelAPIResponse {
	return poolAlibabaServicecenterWorkcardCancelAPIResponse.Get().(*AlibabaServicecenterWorkcardCancelAPIResponse)
}

// ReleaseAlibabaServicecenterWorkcardCancelAPIResponse 将 AlibabaServicecenterWorkcardCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaServicecenterWorkcardCancelAPIResponse(v *AlibabaServicecenterWorkcardCancelAPIResponse) {
	v.Reset()
	poolAlibabaServicecenterWorkcardCancelAPIResponse.Put(v)
}

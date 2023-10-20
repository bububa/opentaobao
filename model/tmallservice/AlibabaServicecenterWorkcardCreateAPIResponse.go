package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterWorkcardCreateAPIResponse 服务平台工单创建接口 API返回值
// alibaba.servicecenter.workcard.create
//
// 创建服务平台工单
type AlibabaServicecenterWorkcardCreateAPIResponse struct {
	model.CommonResponse
	AlibabaServicecenterWorkcardCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaServicecenterWorkcardCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaServicecenterWorkcardCreateAPIResponseModel).Reset()
}

// AlibabaServicecenterWorkcardCreateAPIResponseModel is 服务平台工单创建接口 成功返回结果
type AlibabaServicecenterWorkcardCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_workcard_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaServicecenterWorkcardCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaServicecenterWorkcardCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterWorkcardCreateAPIResponse)
	},
}

// GetAlibabaServicecenterWorkcardCreateAPIResponse 从 sync.Pool 获取 AlibabaServicecenterWorkcardCreateAPIResponse
func GetAlibabaServicecenterWorkcardCreateAPIResponse() *AlibabaServicecenterWorkcardCreateAPIResponse {
	return poolAlibabaServicecenterWorkcardCreateAPIResponse.Get().(*AlibabaServicecenterWorkcardCreateAPIResponse)
}

// ReleaseAlibabaServicecenterWorkcardCreateAPIResponse 将 AlibabaServicecenterWorkcardCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaServicecenterWorkcardCreateAPIResponse(v *AlibabaServicecenterWorkcardCreateAPIResponse) {
	v.Reset()
	poolAlibabaServicecenterWorkcardCreateAPIResponse.Put(v)
}

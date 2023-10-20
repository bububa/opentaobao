package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterFulfiltaskCreateAPIResponse 合单生成核销单 API返回值
// alibaba.servicecenter.fulfiltask.create
//
// 服务对工单进行合单，合单的结果是生成核销单
type AlibabaServicecenterFulfiltaskCreateAPIResponse struct {
	model.CommonResponse
	AlibabaServicecenterFulfiltaskCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaServicecenterFulfiltaskCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaServicecenterFulfiltaskCreateAPIResponseModel).Reset()
}

// AlibabaServicecenterFulfiltaskCreateAPIResponseModel is 合单生成核销单 成功返回结果
type AlibabaServicecenterFulfiltaskCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_fulfiltask_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaServicecenterFulfiltaskCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaServicecenterFulfiltaskCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaServicecenterFulfiltaskCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterFulfiltaskCreateAPIResponse)
	},
}

// GetAlibabaServicecenterFulfiltaskCreateAPIResponse 从 sync.Pool 获取 AlibabaServicecenterFulfiltaskCreateAPIResponse
func GetAlibabaServicecenterFulfiltaskCreateAPIResponse() *AlibabaServicecenterFulfiltaskCreateAPIResponse {
	return poolAlibabaServicecenterFulfiltaskCreateAPIResponse.Get().(*AlibabaServicecenterFulfiltaskCreateAPIResponse)
}

// ReleaseAlibabaServicecenterFulfiltaskCreateAPIResponse 将 AlibabaServicecenterFulfiltaskCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaServicecenterFulfiltaskCreateAPIResponse(v *AlibabaServicecenterFulfiltaskCreateAPIResponse) {
	v.Reset()
	poolAlibabaServicecenterFulfiltaskCreateAPIResponse.Put(v)
}

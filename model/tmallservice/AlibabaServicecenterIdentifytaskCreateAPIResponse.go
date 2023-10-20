package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterIdentifytaskCreateAPIResponse 创建核销单 API返回值
// alibaba.servicecenter.identifytask.create
//
// 创建核销单
type AlibabaServicecenterIdentifytaskCreateAPIResponse struct {
	model.CommonResponse
	AlibabaServicecenterIdentifytaskCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaServicecenterIdentifytaskCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaServicecenterIdentifytaskCreateAPIResponseModel).Reset()
}

// AlibabaServicecenterIdentifytaskCreateAPIResponseModel is 创建核销单 成功返回结果
type AlibabaServicecenterIdentifytaskCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_identifytask_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaServicecenterIdentifytaskCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaServicecenterIdentifytaskCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterIdentifytaskCreateAPIResponse)
	},
}

// GetAlibabaServicecenterIdentifytaskCreateAPIResponse 从 sync.Pool 获取 AlibabaServicecenterIdentifytaskCreateAPIResponse
func GetAlibabaServicecenterIdentifytaskCreateAPIResponse() *AlibabaServicecenterIdentifytaskCreateAPIResponse {
	return poolAlibabaServicecenterIdentifytaskCreateAPIResponse.Get().(*AlibabaServicecenterIdentifytaskCreateAPIResponse)
}

// ReleaseAlibabaServicecenterIdentifytaskCreateAPIResponse 将 AlibabaServicecenterIdentifytaskCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaServicecenterIdentifytaskCreateAPIResponse(v *AlibabaServicecenterIdentifytaskCreateAPIResponse) {
	v.Reset()
	poolAlibabaServicecenterIdentifytaskCreateAPIResponse.Put(v)
}

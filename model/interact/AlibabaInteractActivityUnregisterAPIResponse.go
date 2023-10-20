package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractActivityUnregisterAPIResponse ISV互动应用活动关闭服务 API返回值
// alibaba.interact.activity.unregister
//
// 卖家在ISV互动应用中设置的活动主动关闭的服务
type AlibabaInteractActivityUnregisterAPIResponse struct {
	model.CommonResponse
	AlibabaInteractActivityUnregisterAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractActivityUnregisterAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractActivityUnregisterAPIResponseModel).Reset()
}

// AlibabaInteractActivityUnregisterAPIResponseModel is ISV互动应用活动关闭服务 成功返回结果
type AlibabaInteractActivityUnregisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_activity_unregister_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关闭活动结果
	UnregisterResult *AllsparkResult `json:"unregister_result,omitempty" xml:"unregister_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractActivityUnregisterAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UnregisterResult = nil
}

var poolAlibabaInteractActivityUnregisterAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractActivityUnregisterAPIResponse)
	},
}

// GetAlibabaInteractActivityUnregisterAPIResponse 从 sync.Pool 获取 AlibabaInteractActivityUnregisterAPIResponse
func GetAlibabaInteractActivityUnregisterAPIResponse() *AlibabaInteractActivityUnregisterAPIResponse {
	return poolAlibabaInteractActivityUnregisterAPIResponse.Get().(*AlibabaInteractActivityUnregisterAPIResponse)
}

// ReleaseAlibabaInteractActivityUnregisterAPIResponse 将 AlibabaInteractActivityUnregisterAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractActivityUnregisterAPIResponse(v *AlibabaInteractActivityUnregisterAPIResponse) {
	v.Reset()
	poolAlibabaInteractActivityUnregisterAPIResponse.Put(v)
}

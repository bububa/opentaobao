package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractActivityPushtoalicomAPIResponse 小铺isv推广流量活动到流量聚乐部 API返回值
// alibaba.interact.activity.pushtoalicom
//
// 涉及到流量包的小铺isv，将活动推送到流量聚乐部
type AlibabaInteractActivityPushtoalicomAPIResponse struct {
	model.CommonResponse
	AlibabaInteractActivityPushtoalicomAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractActivityPushtoalicomAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractActivityPushtoalicomAPIResponseModel).Reset()
}

// AlibabaInteractActivityPushtoalicomAPIResponseModel is 小铺isv推广流量活动到流量聚乐部 成功返回结果
type AlibabaInteractActivityPushtoalicomAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_activity_pushtoalicom_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推送成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractActivityPushtoalicomAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaInteractActivityPushtoalicomAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractActivityPushtoalicomAPIResponse)
	},
}

// GetAlibabaInteractActivityPushtoalicomAPIResponse 从 sync.Pool 获取 AlibabaInteractActivityPushtoalicomAPIResponse
func GetAlibabaInteractActivityPushtoalicomAPIResponse() *AlibabaInteractActivityPushtoalicomAPIResponse {
	return poolAlibabaInteractActivityPushtoalicomAPIResponse.Get().(*AlibabaInteractActivityPushtoalicomAPIResponse)
}

// ReleaseAlibabaInteractActivityPushtoalicomAPIResponse 将 AlibabaInteractActivityPushtoalicomAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractActivityPushtoalicomAPIResponse(v *AlibabaInteractActivityPushtoalicomAPIResponse) {
	v.Reset()
	poolAlibabaInteractActivityPushtoalicomAPIResponse.Put(v)
}

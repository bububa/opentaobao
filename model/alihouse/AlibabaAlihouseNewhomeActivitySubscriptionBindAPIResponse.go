package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse 销售活动绑定认购商品 API返回值
// alibaba.alihouse.newhome.activity.subscription.bind
//
// 销售活动绑定认购商品
type AlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponseModel is 销售活动绑定认购商品 成功返回结果
type AlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_activity_subscription_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回报文参数
	Result *AlibabaAlihouseNewhomeActivitySubscriptionBindResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse
func GetAlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse() *AlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse {
	return poolAlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse.Get().(*AlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse 将 AlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse(v *AlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse.Put(v)
}

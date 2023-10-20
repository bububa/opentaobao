package itpolicy

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripItPolicyDeleteAPIResponse 【国际机票销售规则】单条删除 API返回值
// taobao.alitrip.it.policy.delete
//
// 销售规则删除接口，可以根据taobaoId或outId删除，根据outId删除时，如果outId不唯一，返回失败
type TaobaoAlitripItPolicyDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripItPolicyDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripItPolicyDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripItPolicyDeleteAPIResponseModel).Reset()
}

// TaobaoAlitripItPolicyDeleteAPIResponseModel is 【国际机票销售规则】单条删除 成功返回结果
type TaobaoAlitripItPolicyDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_it_policy_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 扩展字段
	ExtendAttributes string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripItPolicyDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ExtendAttributes = ""
}

var poolTaobaoAlitripItPolicyDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripItPolicyDeleteAPIResponse)
	},
}

// GetTaobaoAlitripItPolicyDeleteAPIResponse 从 sync.Pool 获取 TaobaoAlitripItPolicyDeleteAPIResponse
func GetTaobaoAlitripItPolicyDeleteAPIResponse() *TaobaoAlitripItPolicyDeleteAPIResponse {
	return poolTaobaoAlitripItPolicyDeleteAPIResponse.Get().(*TaobaoAlitripItPolicyDeleteAPIResponse)
}

// ReleaseTaobaoAlitripItPolicyDeleteAPIResponse 将 TaobaoAlitripItPolicyDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripItPolicyDeleteAPIResponse(v *TaobaoAlitripItPolicyDeleteAPIResponse) {
	v.Reset()
	poolTaobaoAlitripItPolicyDeleteAPIResponse.Put(v)
}

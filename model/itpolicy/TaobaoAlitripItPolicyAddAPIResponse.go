package itpolicy

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripItPolicyAddAPIResponse 【国际机票销售规则】单条新增 API返回值
// taobao.alitrip.it.policy.add
//
// 销售规则新增，成功返回taobaoId
type TaobaoAlitripItPolicyAddAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripItPolicyAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripItPolicyAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripItPolicyAddAPIResponseModel).Reset()
}

// TaobaoAlitripItPolicyAddAPIResponseModel is 【国际机票销售规则】单条新增 成功返回结果
type TaobaoAlitripItPolicyAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_it_policy_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 扩展字段
	ExtendAttributes string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
	// 淘宝政策id
	TaobaoId int64 `json:"taobao_id,omitempty" xml:"taobao_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripItPolicyAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ExtendAttributes = ""
	m.TaobaoId = 0
}

var poolTaobaoAlitripItPolicyAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripItPolicyAddAPIResponse)
	},
}

// GetTaobaoAlitripItPolicyAddAPIResponse 从 sync.Pool 获取 TaobaoAlitripItPolicyAddAPIResponse
func GetTaobaoAlitripItPolicyAddAPIResponse() *TaobaoAlitripItPolicyAddAPIResponse {
	return poolTaobaoAlitripItPolicyAddAPIResponse.Get().(*TaobaoAlitripItPolicyAddAPIResponse)
}

// ReleaseTaobaoAlitripItPolicyAddAPIResponse 将 TaobaoAlitripItPolicyAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripItPolicyAddAPIResponse(v *TaobaoAlitripItPolicyAddAPIResponse) {
	v.Reset()
	poolTaobaoAlitripItPolicyAddAPIResponse.Put(v)
}

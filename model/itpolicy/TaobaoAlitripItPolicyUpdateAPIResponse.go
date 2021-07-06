package itpolicy

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripItPolicyUpdateAPIResponse 【国际机票销售规则】单条更新 API返回值
// taobao.alitrip.it.policy.update
//
// 销售规则更新接口，可以根据taobaoId或outId修改，outId不唯一时，不能用outId修改。
type TaobaoAlitripItPolicyUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripItPolicyUpdateAPIResponseModel
}

// TaobaoAlitripItPolicyUpdateAPIResponseModel is 【国际机票销售规则】单条更新 成功返回结果
type TaobaoAlitripItPolicyUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_it_policy_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 扩展字段
	ExtendAttributes string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
	// 淘宝政策id
	TaobaoId int64 `json:"taobao_id,omitempty" xml:"taobao_id,omitempty"`
}

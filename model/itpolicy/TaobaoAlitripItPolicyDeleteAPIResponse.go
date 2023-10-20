package itpolicy

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripitpolicydeleteAPIResponse 【国际机票销售规则】单条删除 API返回值
// taobao.alitrip.it.policy.delete
//
// 销售规则删除接口，可以根据taobaoId或outId删除，根据outId删除时，如果outId不唯一，返回失败
type TaobaoalitripitpolicydeleteAPIResponse struct {
	model.CommonResponse
	TaobaoalitripitpolicydeleteAPIResponseModel
}

// TaobaoalitripitpolicydeleteAPIResponseModel is 【国际机票销售规则】单条删除 成功返回结果
type TaobaoalitripitpolicydeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_it_policy_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 扩展字段
	ExtendAttributes string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
}

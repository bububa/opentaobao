package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibababaichuanctguserrelationAPIResponse 用户 API返回值
// alibaba.baichuan.ctg.user.relation
//
// 提供给优酷查询道长和淘宝账户的绑定关系
type AlibababaichuanctguserrelationAPIResponse struct {
	model.CommonResponse
	AlibababaichuanctguserrelationAPIResponseModel
}

// AlibababaichuanctguserrelationAPIResponseModel is 用户 成功返回结果
type AlibababaichuanctguserrelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_baichuan_ctg_user_relation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的整体结果
	Result *AlibababaichuanctguserrelationResult `json:"result,omitempty" xml:"result,omitempty"`
}

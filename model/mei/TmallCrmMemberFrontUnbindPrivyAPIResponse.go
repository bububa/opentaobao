package mei

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallcrmmemberfrontunbindprivyAPIResponse 品牌会员解绑(隐私号版） API返回值
// tmall.crm.member.front.unbind.privy
//
// 品牌会员解绑功能
type TmallcrmmemberfrontunbindprivyAPIResponse struct {
	model.CommonResponse
	TmallcrmmemberfrontunbindprivyAPIResponseModel
}

// TmallcrmmemberfrontunbindprivyAPIResponseModel is 品牌会员解绑(隐私号版） 成功返回结果
type TmallcrmmemberfrontunbindprivyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_crm_member_front_unbind_privy_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口调用是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

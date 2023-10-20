package crm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmmembersgetprivyAPIResponse 获取卖家会员(基本查询) API返回值
// taobao.crm.members.get.privy
//
// 查询卖家的会员，进行基本的查询，返回符合条件的会员列表
type TaobaocrmmembersgetprivyAPIResponse struct {
	model.CommonResponse
	TaobaocrmmembersgetprivyAPIResponseModel
}

// TaobaocrmmembersgetprivyAPIResponseModel is 获取卖家会员(基本查询) 成功返回结果
type TaobaocrmmembersgetprivyAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_members_get_privy_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据一定条件查询到卖家的会员
	Members []BasicMember `json:"members,omitempty" xml:"members>basic_member,omitempty"`
	// 记录总数
	TotalResult int64 `json:"total_result,omitempty" xml:"total_result,omitempty"`
}

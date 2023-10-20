package crm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmmemberssearchprivyAPIResponse 获取卖家会员(高级查询) API返回值
// taobao.crm.members.search.privy
//
// 会员列表的高级查询，接口返回符合条件的会员列表.&lt;br&gt;&lt;br/&gt;注：建议获取09年以后的数据，09年之前的数据不是很完整
type TaobaocrmmemberssearchprivyAPIResponse struct {
	model.CommonResponse
	TaobaocrmmemberssearchprivyAPIResponseModel
}

// TaobaocrmmemberssearchprivyAPIResponseModel is 获取卖家会员(高级查询) 成功返回结果
type TaobaocrmmemberssearchprivyAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_members_search_privy_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据一定条件查询的卖家会员
	Members []CrmMember `json:"members,omitempty" xml:"members>crm_member,omitempty"`
	// 记录的总条数
	TotalResult int64 `json:"total_result,omitempty" xml:"total_result,omitempty"`
}

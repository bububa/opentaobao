package crm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmMemberGroupGetAPIResponse
获取买家身上的标签 API返回值
taobao.crm.member.group.get

获取买家身上的标签，不返回标签的总人数 */
type TaobaoCrmMemberGroupGetAPIResponse struct {
	model.CommonResponse
	TaobaoCrmMemberGroupGetAPIResponseModel
}

// TaobaoCrmMemberGroupGetAPIResponseModel is 获取买家身上的标签 成功返回结果
type TaobaoCrmMemberGroupGetAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_member_group_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询到的当前卖家的当前页的会员
	Groups []Group `json:"groups,omitempty" xml:"groups>group,omitempty"`
}

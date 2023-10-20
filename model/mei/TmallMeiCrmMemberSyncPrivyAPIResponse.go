package mei

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallmeicrmmembersyncprivyAPIResponse 同步推送会员信息(隐私号版本) API返回值
// tmall.mei.crm.member.sync.privy
//
// 品牌方通过该api主动推送会员信息。使用场景包括
// 1.用户在线上注册后，线下补充信息后，同步到线上。
// 2.其他情况的主动推送变更。
type TmallmeicrmmembersyncprivyAPIResponse struct {
	model.CommonResponse
	TmallmeicrmmembersyncprivyAPIResponseModel
}

// TmallmeicrmmembersyncprivyAPIResponseModel is 同步推送会员信息(隐私号版本) 成功返回结果
type TmallmeicrmmembersyncprivyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_mei_crm_member_sync_privy_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 处理的其他信息
	MeiExtraInfo string `json:"mei_extra_info,omitempty" xml:"mei_extra_info,omitempty"`
}

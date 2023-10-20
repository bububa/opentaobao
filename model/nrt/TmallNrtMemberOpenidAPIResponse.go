package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtMemberOpenidAPIResponse 根据会员手机查询openId API返回值
// tmall.nrt.member.openid
//
// 根据会员手机查询openId
type TmallNrtMemberOpenidAPIResponse struct {
	model.CommonResponse
	TmallNrtMemberOpenidAPIResponseModel
}

// TmallNrtMemberOpenidAPIResponseModel is 根据会员手机查询openId 成功返回结果
type TmallNrtMemberOpenidAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_member_openid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TmallNrtMemberOpenidResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

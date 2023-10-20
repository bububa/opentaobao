package xhotelcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelMemberDerbyStateSyncAPIResponse 德比侧同步卡、券状态接口 API返回值
// taobao.xhotel.member.derby.state.sync
//
// 德比侧同步卡、券状态接口
type TaobaoXhotelMemberDerbyStateSyncAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelMemberDerbyStateSyncAPIResponseModel
}

// TaobaoXhotelMemberDerbyStateSyncAPIResponseModel is 德比侧同步卡、券状态接口 成功返回结果
type TaobaoXhotelMemberDerbyStateSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_member_derby_state_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *MsdResult `json:"result,omitempty" xml:"result,omitempty"`
}

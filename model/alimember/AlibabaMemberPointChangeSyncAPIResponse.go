package alimember

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberPointChangeSyncAPIResponse 成长值/积分变更记录同步 API返回值
// alibaba.member.point.change.sync
//
// 成长值/积分变更记录同步
type AlibabaMemberPointChangeSyncAPIResponse struct {
	model.CommonResponse
	AlibabaMemberPointChangeSyncAPIResponseModel
}

// AlibabaMemberPointChangeSyncAPIResponseModel is 成长值/积分变更记录同步 成功返回结果
type AlibabaMemberPointChangeSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_member_point_change_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码 ，比如 E1003 用户不是会员
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

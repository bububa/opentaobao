package alimember

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberPointOperateAPIResponse 积分操作 API返回值
// alibaba.member.point.operate
//
// 消费会员积分
type AlibabaMemberPointOperateAPIResponse struct {
	model.CommonResponse
	AlibabaMemberPointOperateAPIResponseModel
}

// AlibabaMemberPointOperateAPIResponseModel is 积分操作 成功返回结果
type AlibabaMemberPointOperateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_member_point_operate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaMemberPointOperateTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

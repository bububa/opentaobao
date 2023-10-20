package legalsuit

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalsuitcourtafterpushAPIResponse 更新或者新增庭后信息 API返回值
// alibaba.legal.suit.court.after.push
//
// 供外部ISV供应商 推送庭后信息给集团诉讼
type AlibabalegalsuitcourtafterpushAPIResponse struct {
	model.CommonResponse
	AlibabalegalsuitcourtafterpushAPIResponseModel
}

// AlibabalegalsuitcourtafterpushAPIResponseModel is 更新或者新增庭后信息 成功返回结果
type AlibabalegalsuitcourtafterpushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_suit_court_after_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaiworkmcmsgsendmobileAPIResponse 发送消息给手机用户 API返回值
// alibaba.iwork.mc.msg.sendmobile
//
// 给手机用户发送对应操作结果的消息
type AlibabaiworkmcmsgsendmobileAPIResponse struct {
	model.CommonResponse
	AlibabaiworkmcmsgsendmobileAPIResponseModel
}

// AlibabaiworkmcmsgsendmobileAPIResponseModel is 发送消息给手机用户 成功返回结果
type AlibabaiworkmcmsgsendmobileAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_iwork_mc_msg_sendmobile_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

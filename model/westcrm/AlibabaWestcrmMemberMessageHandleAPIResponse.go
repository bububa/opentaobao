package westcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmMemberMessageHandleAPIResponse
处理Usc会员消息接口 API返回值
alibaba.westcrm.member.message.handle

处理Usc会员消息接口 */
type AlibabaWestcrmMemberMessageHandleAPIResponse struct {
	model.CommonResponse
	AlibabaWestcrmMemberMessageHandleAPIResponseModel
}

// AlibabaWestcrmMemberMessageHandleAPIResponseModel is 处理Usc会员消息接口 成功返回结果
type AlibabaWestcrmMemberMessageHandleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_westcrm_member_message_handle_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

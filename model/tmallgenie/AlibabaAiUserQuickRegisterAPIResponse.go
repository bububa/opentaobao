package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAiUserQuickRegisterAPIResponse 精灵用户注册申请 API返回值
// alibaba.ai.user.quick.register
//
// 人工智能实验室精灵用户注册申请接口，开放给Iot厂商做厂商会员数据上报
type AlibabaAiUserQuickRegisterAPIResponse struct {
	model.CommonResponse
	AlibabaAiUserQuickRegisterAPIResponseModel
}

// AlibabaAiUserQuickRegisterAPIResponseModel is 精灵用户注册申请 成功返回结果
type AlibabaAiUserQuickRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ai_user_quick_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 扩展对象，保留字段
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 结果详细描述
	MessageDetail string `json:"message_detail,omitempty" xml:"message_detail,omitempty"`
	// 结果码
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}

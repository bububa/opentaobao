package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIworkMcMsgSendmobileAPIResponse 发送消息给手机用户 API返回值
// alibaba.iwork.mc.msg.sendmobile
//
// 给手机用户发送对应操作结果的消息
type AlibabaIworkMcMsgSendmobileAPIResponse struct {
	model.CommonResponse
	AlibabaIworkMcMsgSendmobileAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIworkMcMsgSendmobileAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIworkMcMsgSendmobileAPIResponseModel).Reset()
}

// AlibabaIworkMcMsgSendmobileAPIResponseModel is 发送消息给手机用户 成功返回结果
type AlibabaIworkMcMsgSendmobileAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_iwork_mc_msg_sendmobile_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIworkMcMsgSendmobileAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIworkMcMsgSendmobileAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIworkMcMsgSendmobileAPIResponse)
	},
}

// GetAlibabaIworkMcMsgSendmobileAPIResponse 从 sync.Pool 获取 AlibabaIworkMcMsgSendmobileAPIResponse
func GetAlibabaIworkMcMsgSendmobileAPIResponse() *AlibabaIworkMcMsgSendmobileAPIResponse {
	return poolAlibabaIworkMcMsgSendmobileAPIResponse.Get().(*AlibabaIworkMcMsgSendmobileAPIResponse)
}

// ReleaseAlibabaIworkMcMsgSendmobileAPIResponse 将 AlibabaIworkMcMsgSendmobileAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIworkMcMsgSendmobileAPIResponse(v *AlibabaIworkMcMsgSendmobileAPIResponse) {
	v.Reset()
	poolAlibabaIworkMcMsgSendmobileAPIResponse.Put(v)
}

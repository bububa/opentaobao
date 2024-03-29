package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIworkMcMsgSenddefaultAPIResponse 给注册用户发送消息 API返回值
// alibaba.iwork.mc.msg.senddefault
//
// 给神鲸注册用户发送对应操作结果的消息
type AlibabaIworkMcMsgSenddefaultAPIResponse struct {
	model.CommonResponse
	AlibabaIworkMcMsgSenddefaultAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIworkMcMsgSenddefaultAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIworkMcMsgSenddefaultAPIResponseModel).Reset()
}

// AlibabaIworkMcMsgSenddefaultAPIResponseModel is 给注册用户发送消息 成功返回结果
type AlibabaIworkMcMsgSenddefaultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_iwork_mc_msg_senddefault_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIworkMcMsgSenddefaultAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIworkMcMsgSenddefaultAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIworkMcMsgSenddefaultAPIResponse)
	},
}

// GetAlibabaIworkMcMsgSenddefaultAPIResponse 从 sync.Pool 获取 AlibabaIworkMcMsgSenddefaultAPIResponse
func GetAlibabaIworkMcMsgSenddefaultAPIResponse() *AlibabaIworkMcMsgSenddefaultAPIResponse {
	return poolAlibabaIworkMcMsgSenddefaultAPIResponse.Get().(*AlibabaIworkMcMsgSenddefaultAPIResponse)
}

// ReleaseAlibabaIworkMcMsgSenddefaultAPIResponse 将 AlibabaIworkMcMsgSenddefaultAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIworkMcMsgSenddefaultAPIResponse(v *AlibabaIworkMcMsgSenddefaultAPIResponse) {
	v.Reset()
	poolAlibabaIworkMcMsgSenddefaultAPIResponse.Put(v)
}

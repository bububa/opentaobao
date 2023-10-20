package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMessageaccountMesssageNormalSendAPIResponse 下行普通消息 API返回值
// taobao.messageaccount.messsage.normal.send
//
// 消息号下行单个普通消息
type TaobaoMessageaccountMesssageNormalSendAPIResponse struct {
	model.CommonResponse
	TaobaoMessageaccountMesssageNormalSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMessageaccountMesssageNormalSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMessageaccountMesssageNormalSendAPIResponseModel).Reset()
}

// TaobaoMessageaccountMesssageNormalSendAPIResponseModel is 下行普通消息 成功返回结果
type TaobaoMessageaccountMesssageNormalSendAPIResponseModel struct {
	XMLName xml.Name `xml:"messageaccount_messsage_normal_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMessageaccountMesssageNormalSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
}

var poolTaobaoMessageaccountMesssageNormalSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMessageaccountMesssageNormalSendAPIResponse)
	},
}

// GetTaobaoMessageaccountMesssageNormalSendAPIResponse 从 sync.Pool 获取 TaobaoMessageaccountMesssageNormalSendAPIResponse
func GetTaobaoMessageaccountMesssageNormalSendAPIResponse() *TaobaoMessageaccountMesssageNormalSendAPIResponse {
	return poolTaobaoMessageaccountMesssageNormalSendAPIResponse.Get().(*TaobaoMessageaccountMesssageNormalSendAPIResponse)
}

// ReleaseTaobaoMessageaccountMesssageNormalSendAPIResponse 将 TaobaoMessageaccountMesssageNormalSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMessageaccountMesssageNormalSendAPIResponse(v *TaobaoMessageaccountMesssageNormalSendAPIResponse) {
	v.Reset()
	poolTaobaoMessageaccountMesssageNormalSendAPIResponse.Put(v)
}

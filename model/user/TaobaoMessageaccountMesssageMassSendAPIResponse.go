package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMessageaccountMesssageMassSendAPIResponse 消息号开放-消息群发 API返回值
// taobao.messageaccount.messsage.mass.send
//
// 外部 isv 调用该进口来进行消息号消息的群发
type TaobaoMessageaccountMesssageMassSendAPIResponse struct {
	model.CommonResponse
	TaobaoMessageaccountMesssageMassSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMessageaccountMesssageMassSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMessageaccountMesssageMassSendAPIResponseModel).Reset()
}

// TaobaoMessageaccountMesssageMassSendAPIResponseModel is 消息号开放-消息群发 成功返回结果
type TaobaoMessageaccountMesssageMassSendAPIResponseModel struct {
	XMLName xml.Name `xml:"messageaccount_messsage_mass_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoMessageaccountMesssageMassSendResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMessageaccountMesssageMassSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoMessageaccountMesssageMassSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMessageaccountMesssageMassSendAPIResponse)
	},
}

// GetTaobaoMessageaccountMesssageMassSendAPIResponse 从 sync.Pool 获取 TaobaoMessageaccountMesssageMassSendAPIResponse
func GetTaobaoMessageaccountMesssageMassSendAPIResponse() *TaobaoMessageaccountMesssageMassSendAPIResponse {
	return poolTaobaoMessageaccountMesssageMassSendAPIResponse.Get().(*TaobaoMessageaccountMesssageMassSendAPIResponse)
}

// ReleaseTaobaoMessageaccountMesssageMassSendAPIResponse 将 TaobaoMessageaccountMesssageMassSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMessageaccountMesssageMassSendAPIResponse(v *TaobaoMessageaccountMesssageMassSendAPIResponse) {
	v.Reset()
	poolTaobaoMessageaccountMesssageMassSendAPIResponse.Put(v)
}

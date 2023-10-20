package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopOaidDecryptAPIResponse OAID解密 API返回值
// taobao.top.oaid.decrypt
//
// 解码OAID(Open Addressee ID)，返回收件人信息。
type TaobaoTopOaidDecryptAPIResponse struct {
	model.CommonResponse
	TaobaoTopOaidDecryptAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTopOaidDecryptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTopOaidDecryptAPIResponseModel).Reset()
}

// TaobaoTopOaidDecryptAPIResponseModel is OAID解密 成功返回结果
type TaobaoTopOaidDecryptAPIResponseModel struct {
	XMLName xml.Name `xml:"top_oaid_decrypt_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 收件人列表
	ReceiverList []Receiver `json:"receiver_list,omitempty" xml:"receiver_list>receiver,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTopOaidDecryptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ReceiverList = m.ReceiverList[:0]
}

var poolTaobaoTopOaidDecryptAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTopOaidDecryptAPIResponse)
	},
}

// GetTaobaoTopOaidDecryptAPIResponse 从 sync.Pool 获取 TaobaoTopOaidDecryptAPIResponse
func GetTaobaoTopOaidDecryptAPIResponse() *TaobaoTopOaidDecryptAPIResponse {
	return poolTaobaoTopOaidDecryptAPIResponse.Get().(*TaobaoTopOaidDecryptAPIResponse)
}

// ReleaseTaobaoTopOaidDecryptAPIResponse 将 TaobaoTopOaidDecryptAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTopOaidDecryptAPIResponse(v *TaobaoTopOaidDecryptAPIResponse) {
	v.Reset()
	poolTaobaoTopOaidDecryptAPIResponse.Put(v)
}

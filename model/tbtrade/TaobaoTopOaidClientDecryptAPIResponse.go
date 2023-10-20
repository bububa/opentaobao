package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopOaidClientDecryptAPIResponse 端侧OAID解密 API返回值
// taobao.top.oaid.client.decrypt
//
// 解码OAID(Open Addressee ID)，返回收件人信息。该接口用于客户端直接查看订单隐私数据，解密数据不经过ISV服务器，且包含风控等安全检测。
type TaobaoTopOaidClientDecryptAPIResponse struct {
	model.CommonResponse
	TaobaoTopOaidClientDecryptAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTopOaidClientDecryptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTopOaidClientDecryptAPIResponseModel).Reset()
}

// TaobaoTopOaidClientDecryptAPIResponseModel is 端侧OAID解密 成功返回结果
type TaobaoTopOaidClientDecryptAPIResponseModel struct {
	XMLName xml.Name `xml:"top_oaid_client_decrypt_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 收件人列表
	ReceiverList []Receiver `json:"receiver_list,omitempty" xml:"receiver_list>receiver,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTopOaidClientDecryptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ReceiverList = m.ReceiverList[:0]
}

var poolTaobaoTopOaidClientDecryptAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTopOaidClientDecryptAPIResponse)
	},
}

// GetTaobaoTopOaidClientDecryptAPIResponse 从 sync.Pool 获取 TaobaoTopOaidClientDecryptAPIResponse
func GetTaobaoTopOaidClientDecryptAPIResponse() *TaobaoTopOaidClientDecryptAPIResponse {
	return poolTaobaoTopOaidClientDecryptAPIResponse.Get().(*TaobaoTopOaidClientDecryptAPIResponse)
}

// ReleaseTaobaoTopOaidClientDecryptAPIResponse 将 TaobaoTopOaidClientDecryptAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTopOaidClientDecryptAPIResponse(v *TaobaoTopOaidClientDecryptAPIResponse) {
	v.Reset()
	poolTaobaoTopOaidClientDecryptAPIResponse.Put(v)
}

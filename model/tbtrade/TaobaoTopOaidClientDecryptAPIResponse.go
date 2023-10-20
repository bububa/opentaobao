package tbtrade

import (
	"encoding/xml"

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

// TaobaoTopOaidClientDecryptAPIResponseModel is 端侧OAID解密 成功返回结果
type TaobaoTopOaidClientDecryptAPIResponseModel struct {
	XMLName xml.Name `xml:"top_oaid_client_decrypt_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 收件人列表
	ReceiverList []Receiver `json:"receiver_list,omitempty" xml:"receiver_list>receiver,omitempty"`
}

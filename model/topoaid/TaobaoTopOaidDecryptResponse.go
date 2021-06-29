package topoaid

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
OAID解密 APIResponse
taobao.top.oaid.decrypt

解码OAID(Open Addressee ID)，返回收件人信息。
*/
type TaobaoTopOaidDecryptAPIResponse struct {
    model.CommonResponse
    TaobaoTopOaidDecryptResponse
}

type TaobaoTopOaidDecryptResponse struct {
    XMLName xml.Name `xml:"top_oaid_decrypt_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 收件人列表
    
    ReceiverList   []Receiver `json:"receiver_list,omitempty" xml:"receiver_list>receiver,omitempty"`
    
    
}

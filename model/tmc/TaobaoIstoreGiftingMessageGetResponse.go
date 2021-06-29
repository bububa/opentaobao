package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
gifting消息获取 APIResponse
taobao.istore.gifting.message.get

该api通过参数查询对应的gifting消息
*/
type TaobaoIstoreGiftingMessageGetAPIResponse struct {
    model.CommonResponse
    TaobaoIstoreGiftingMessageGetResponse
}

type TaobaoIstoreGiftingMessageGetResponse struct {
    XMLName xml.Name `xml:"istore_gifting_message_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoIstoreGiftingMessageGetResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}

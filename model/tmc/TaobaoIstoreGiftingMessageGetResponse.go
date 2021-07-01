package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
gifting消息获取 API返回值 
taobao.istore.gifting.message.get

该api通过参数查询对应的gifting消息
*/
type TaobaoIstoreGiftingMessageGetAPIResponse struct {
    model.CommonResponse
    TaobaoIstoreGiftingMessageGetResponse
}

// gifting消息获取 成功返回结果
type TaobaoIstoreGiftingMessageGetResponse struct {
    XMLName xml.Name `xml:"istore_gifting_message_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoIstoreGiftingMessageGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

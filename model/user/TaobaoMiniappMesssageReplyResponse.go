package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
轻店铺下行回复接口 APIResponse
taobao.miniapp.messsage.reply

外部 isv 调用该进口来进行轻店铺消息的回复
*/
type TaobaoMiniappMesssageReplyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"miniapp_messsage_reply_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoMiniappMesssageReplyResult `json:"result,omitempty" xml:"
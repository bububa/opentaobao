package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
下行普通消息 API返回值 
taobao.messageaccount.messsage.normal.send

消息号下行单个普通消息
*/
type TaobaoMessageaccountMesssageNormalSendAPIResponse struct {
    model.CommonResponse
    TaobaoMessageaccountMesssageNormalSendAPIResponseModel
}

// 下行普通消息 成功返回结果
type TaobaoMessageaccountMesssageNormalSendAPIResponseModel struct {
    XMLName xml.Name `xml:"messageaccount_messsage_normal_send_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // model
    Model   string `json:"model,omitempty" xml:"model,omitempty"`
}

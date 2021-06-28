package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
轻店铺下行普通消息给用户 APIResponse
taobao.miniapp.messsage.normal.send

小程序下行单个普通消息
*/
type TaobaoMiniappMesssageNormalSendAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"miniapp_messsage_normal_send_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回消息Id
    
    Model   string `json:"model,omitempty" xml:"
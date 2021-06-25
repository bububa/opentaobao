package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
轻店铺下行普通消息给用户 APIResponse
taobao.miniapp.messsage.normal.send

小程序下行单个普通消息
*/
type TaobaoMiniappMesssageNormalSendAPIResponse struct {
    model.CommonResponse
    Response *TaobaoMiniappMesssageNormalSendResponse `json:"taobao_miniapp_messsage_normal_send_response,omitempty"`
}

type TaobaoMiniappMesssageNormalSendResponse struct {

    // 返回消息Id
    Model   string `json:"model,omitempty"`

}

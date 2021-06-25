package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川发送注册验证码 APIResponse
taobao.baichuan.openaccount.registercode.send

百川发送注册验证码
*/
type TaobaoBaichuanOpenaccountRegistercodeSendAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBaichuanOpenaccountRegistercodeSendResponse `json:"taobao_baichuan_openaccount_registercode_send_response,omitempty"`
}

type TaobaoBaichuanOpenaccountRegistercodeSendResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}

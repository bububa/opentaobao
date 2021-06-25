package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川账号注册 APIResponse
taobao.baichuan.openaccount.register

百川账号注册
*/
type TaobaoBaichuanOpenaccountRegisterAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBaichuanOpenaccountRegisterResponse `json:"taobao_baichuan_openaccount_register_response,omitempty"`
}

type TaobaoBaichuanOpenaccountRegisterResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}

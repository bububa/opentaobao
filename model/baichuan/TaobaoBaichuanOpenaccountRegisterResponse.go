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
    // Response *TaobaoBaichuanOpenaccountRegisterResponse `json:"baichuan_openaccount_register_response,omitempty"` 
    TaobaoBaichuanOpenaccountRegisterResponse
}

/* model for simplify = false
type TaobaoBaichuanOpenaccountRegisterResponse struct {

    // name
    
    Name   string `json:"name,omitempty"`
    

}
*/

type TaobaoBaichuanOpenaccountRegisterResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}

package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川检查注册验证码 APIResponse
taobao.baichuan.openaccount.registercode.check

百川检查注册验证码
*/
type TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBaichuanOpenaccountRegistercodeCheckResponse `json:"baichuan_openaccount_registercode_check_response,omitempty"` 
    TaobaoBaichuanOpenaccountRegistercodeCheckResponse
}

/* model for simplify = false
type TaobaoBaichuanOpenaccountRegistercodeCheckResponse struct {

    // name
    
    Name   string `json:"name,omitempty"`
    

}
*/

type TaobaoBaichuanOpenaccountRegistercodeCheckResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}

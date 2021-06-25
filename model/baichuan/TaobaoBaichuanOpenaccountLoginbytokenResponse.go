package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川TOKEN 登录 APIResponse
taobao.baichuan.openaccount.loginbytoken

百川TOKEN 登录
*/
type TaobaoBaichuanOpenaccountLoginbytokenAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBaichuanOpenaccountLoginbytokenResponse `json:"taobao_baichuan_openaccount_loginbytoken_response,omitempty"`
}

type TaobaoBaichuanOpenaccountLoginbytokenResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}

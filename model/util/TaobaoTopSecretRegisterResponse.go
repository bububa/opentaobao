package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
注册加密账号 APIResponse
taobao.top.secret.register

提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密
*/
type TaobaoTopSecretRegisterAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTopSecretRegisterResponse `json:"taobao_top_secret_register_response,omitempty"`
}

type TaobaoTopSecretRegisterResponse struct {

    // 返回操作是否成功
    Result   bool `json:"result,omitempty"`

}

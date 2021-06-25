package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
注册加密账号 APIRequest
taobao.top.secret.register

提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密
*/
type TaobaoTopSecretRegisterRequest struct {
    model.Params

    // 用户id，保证唯一
    userId   int64 

}

func NewTaobaoTopSecretRegisterRequest() *TaobaoTopSecretRegisterRequest{
    return &TaobaoTopSecretRegisterRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTopSecretRegisterRequest) GetApiMethodName() string {
    return "taobao.top.secret.register"
}

func (r TaobaoTopSecretRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTopSecretRegisterRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoTopSecretRegisterRequest) GetUserId() int64 {
    return r.userId
}


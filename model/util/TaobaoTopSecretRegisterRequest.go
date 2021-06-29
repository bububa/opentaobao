package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
注册加密账号 API请求
taobao.top.secret.register

提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密
*/
type TaobaoTopSecretRegisterRequest struct {
    model.Params
    // 用户id，保证唯一
    userId   int64
}

// 初始化TaobaoTopSecretRegisterRequest对象
func NewTaobaoTopSecretRegisterRequest() *TaobaoTopSecretRegisterRequest{
    return &TaobaoTopSecretRegisterRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTopSecretRegisterRequest) GetApiMethodName() string {
    return "taobao.top.secret.register"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTopSecretRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户id，保证唯一
func (r *TaobaoTopSecretRegisterRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoTopSecretRegisterRequest) GetUserId() int64 {
    return r.userId
}

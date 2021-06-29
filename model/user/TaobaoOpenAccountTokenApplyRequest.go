package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
申请免登Open Account Token API请求
taobao.open.account.token.apply

申请免登Open Account Token
*/
type TaobaoOpenAccountTokenApplyRequest struct {
    model.Params
    // 时间戳单位是毫秒
    tokenTimestamp   int64
    // open account id
    openAccountId   int64
    // isv自己账号的唯一id
    isvAccountId   string
    // 用于防重放的唯一id
    uuid   string
    // ISV APP的登录态时长单位是秒
    loginStateExpireIn   int64
    // 用于透传一些业务附加参数
    ext   string
}

// 初始化TaobaoOpenAccountTokenApplyRequest对象
func NewTaobaoOpenAccountTokenApplyRequest() *TaobaoOpenAccountTokenApplyRequest{
    return &TaobaoOpenAccountTokenApplyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountTokenApplyRequest) GetApiMethodName() string {
    return "taobao.open.account.token.apply"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountTokenApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TokenTimestamp Setter
// 时间戳单位是毫秒
func (r *TaobaoOpenAccountTokenApplyRequest) SetTokenTimestamp(tokenTimestamp int64) error {
    r.tokenTimestamp = tokenTimestamp
    r.Set("token_timestamp", tokenTimestamp)
    return nil
}

// TokenTimestamp Getter
func (r TaobaoOpenAccountTokenApplyRequest) GetTokenTimestamp() int64 {
    return r.tokenTimestamp
}
// OpenAccountId Setter
// open account id
func (r *TaobaoOpenAccountTokenApplyRequest) SetOpenAccountId(openAccountId int64) error {
    r.openAccountId = openAccountId
    r.Set("open_account_id", openAccountId)
    return nil
}

// OpenAccountId Getter
func (r TaobaoOpenAccountTokenApplyRequest) GetOpenAccountId() int64 {
    return r.openAccountId
}
// IsvAccountId Setter
// isv自己账号的唯一id
func (r *TaobaoOpenAccountTokenApplyRequest) SetIsvAccountId(isvAccountId string) error {
    r.isvAccountId = isvAccountId
    r.Set("isv_account_id", isvAccountId)
    return nil
}

// IsvAccountId Getter
func (r TaobaoOpenAccountTokenApplyRequest) GetIsvAccountId() string {
    return r.isvAccountId
}
// Uuid Setter
// 用于防重放的唯一id
func (r *TaobaoOpenAccountTokenApplyRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r TaobaoOpenAccountTokenApplyRequest) GetUuid() string {
    return r.uuid
}
// LoginStateExpireIn Setter
// ISV APP的登录态时长单位是秒
func (r *TaobaoOpenAccountTokenApplyRequest) SetLoginStateExpireIn(loginStateExpireIn int64) error {
    r.loginStateExpireIn = loginStateExpireIn
    r.Set("login_state_expire_in", loginStateExpireIn)
    return nil
}

// LoginStateExpireIn Getter
func (r TaobaoOpenAccountTokenApplyRequest) GetLoginStateExpireIn() int64 {
    return r.loginStateExpireIn
}
// Ext Setter
// 用于透传一些业务附加参数
func (r *TaobaoOpenAccountTokenApplyRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoOpenAccountTokenApplyRequest) GetExt() string {
    return r.ext
}

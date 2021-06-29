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
    _tokenTimestamp   int64
    // open account id
    _openAccountId   int64
    // isv自己账号的唯一id
    _isvAccountId   string
    // 用于防重放的唯一id
    _uuid   string
    // ISV APP的登录态时长单位是秒
    _loginStateExpireIn   int64
    // 用于透传一些业务附加参数
    _ext   string
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
func (r *TaobaoOpenAccountTokenApplyRequest) SetTokenTimestamp(_tokenTimestamp int64) error {
    r._tokenTimestamp = _tokenTimestamp
    r.Set("token_timestamp", _tokenTimestamp)
    return nil
}

// TokenTimestamp Getter
func (r TaobaoOpenAccountTokenApplyRequest) GetTokenTimestamp() int64 {
    return r._tokenTimestamp
}
// OpenAccountId Setter
// open account id
func (r *TaobaoOpenAccountTokenApplyRequest) SetOpenAccountId(_openAccountId int64) error {
    r._openAccountId = _openAccountId
    r.Set("open_account_id", _openAccountId)
    return nil
}

// OpenAccountId Getter
func (r TaobaoOpenAccountTokenApplyRequest) GetOpenAccountId() int64 {
    return r._openAccountId
}
// IsvAccountId Setter
// isv自己账号的唯一id
func (r *TaobaoOpenAccountTokenApplyRequest) SetIsvAccountId(_isvAccountId string) error {
    r._isvAccountId = _isvAccountId
    r.Set("isv_account_id", _isvAccountId)
    return nil
}

// IsvAccountId Getter
func (r TaobaoOpenAccountTokenApplyRequest) GetIsvAccountId() string {
    return r._isvAccountId
}
// Uuid Setter
// 用于防重放的唯一id
func (r *TaobaoOpenAccountTokenApplyRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r TaobaoOpenAccountTokenApplyRequest) GetUuid() string {
    return r._uuid
}
// LoginStateExpireIn Setter
// ISV APP的登录态时长单位是秒
func (r *TaobaoOpenAccountTokenApplyRequest) SetLoginStateExpireIn(_loginStateExpireIn int64) error {
    r._loginStateExpireIn = _loginStateExpireIn
    r.Set("login_state_expire_in", _loginStateExpireIn)
    return nil
}

// LoginStateExpireIn Getter
func (r TaobaoOpenAccountTokenApplyRequest) GetLoginStateExpireIn() int64 {
    return r._loginStateExpireIn
}
// Ext Setter
// 用于透传一些业务附加参数
func (r *TaobaoOpenAccountTokenApplyRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoOpenAccountTokenApplyRequest) GetExt() string {
    return r._ext
}

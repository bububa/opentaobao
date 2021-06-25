package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
申请免登Open Account Token APIRequest
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

func NewTaobaoOpenAccountTokenApplyRequest() *TaobaoOpenAccountTokenApplyRequest{
    return &TaobaoOpenAccountTokenApplyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenAccountTokenApplyRequest) GetApiMethodName() string {
    return "taobao.open.account.token.apply"
}

func (r TaobaoOpenAccountTokenApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenAccountTokenApplyRequest) SetTokenTimestamp(tokenTimestamp int64) error {
    r.tokenTimestamp = tokenTimestamp
    r.Set("token_timestamp", tokenTimestamp)
    return nil
}

func (r TaobaoOpenAccountTokenApplyRequest) GetTokenTimestamp() int64 {
    return r.tokenTimestamp
}

func (r *TaobaoOpenAccountTokenApplyRequest) SetOpenAccountId(openAccountId int64) error {
    r.openAccountId = openAccountId
    r.Set("open_account_id", openAccountId)
    return nil
}

func (r TaobaoOpenAccountTokenApplyRequest) GetOpenAccountId() int64 {
    return r.openAccountId
}

func (r *TaobaoOpenAccountTokenApplyRequest) SetIsvAccountId(isvAccountId string) error {
    r.isvAccountId = isvAccountId
    r.Set("isv_account_id", isvAccountId)
    return nil
}

func (r TaobaoOpenAccountTokenApplyRequest) GetIsvAccountId() string {
    return r.isvAccountId
}

func (r *TaobaoOpenAccountTokenApplyRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r TaobaoOpenAccountTokenApplyRequest) GetUuid() string {
    return r.uuid
}

func (r *TaobaoOpenAccountTokenApplyRequest) SetLoginStateExpireIn(loginStateExpireIn int64) error {
    r.loginStateExpireIn = loginStateExpireIn
    r.Set("login_state_expire_in", loginStateExpireIn)
    return nil
}

func (r TaobaoOpenAccountTokenApplyRequest) GetLoginStateExpireIn() int64 {
    return r.loginStateExpireIn
}

func (r *TaobaoOpenAccountTokenApplyRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoOpenAccountTokenApplyRequest) GetExt() string {
    return r.ext
}


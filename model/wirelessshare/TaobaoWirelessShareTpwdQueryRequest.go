package wirelessshare

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询解析淘口令 APIRequest
taobao.wireless.share.tpwd.query

查询解析淘口令
*/
type TaobaoWirelessShareTpwdQueryRequest struct {
    model.Params

    // 淘口令
    passwordContent   string 

}

func NewTaobaoWirelessShareTpwdQueryRequest() *TaobaoWirelessShareTpwdQueryRequest{
    return &TaobaoWirelessShareTpwdQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWirelessShareTpwdQueryRequest) GetApiMethodName() string {
    return "taobao.wireless.share.tpwd.query"
}

func (r TaobaoWirelessShareTpwdQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWirelessShareTpwdQueryRequest) SetPasswordContent(passwordContent string) error {
    r.passwordContent = passwordContent
    r.Set("password_content", passwordContent)
    return nil
}

func (r TaobaoWirelessShareTpwdQueryRequest) GetPasswordContent() string {
    return r.passwordContent
}


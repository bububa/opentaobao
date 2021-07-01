package wirelessshare

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询解析淘口令 API请求
taobao.wireless.share.tpwd.query

查询解析淘口令
*/
type TaobaoWirelessShareTpwdQueryAPIRequest struct {
    model.Params
    // 淘口令
    _passwordContent   string
}

// 初始化TaobaoWirelessShareTpwdQueryAPIRequest对象
func NewTaobaoWirelessShareTpwdQueryRequest() *TaobaoWirelessShareTpwdQueryAPIRequest{
    return &TaobaoWirelessShareTpwdQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWirelessShareTpwdQueryAPIRequest) GetApiMethodName() string {
    return "taobao.wireless.share.tpwd.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWirelessShareTpwdQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PasswordContent Setter
// 淘口令
func (r *TaobaoWirelessShareTpwdQueryAPIRequest) SetPasswordContent(_passwordContent string) error {
    r._passwordContent = _passwordContent
    r.Set("password_content", _passwordContent)
    return nil
}

// PasswordContent Getter
func (r TaobaoWirelessShareTpwdQueryAPIRequest) GetPasswordContent() string {
    return r._passwordContent
}

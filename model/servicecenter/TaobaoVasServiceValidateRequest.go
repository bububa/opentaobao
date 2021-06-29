package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增值服务订购服务验证 API请求
taobao.vas.service.validate

增值服务订购服务验证
*/
type TaobaoVasServiceValidateRequest struct {
    model.Params
    // 服务编码
    servCode   string
    // 用户昵称
    nick   string
}

// 初始化TaobaoVasServiceValidateRequest对象
func NewTaobaoVasServiceValidateRequest() *TaobaoVasServiceValidateRequest{
    return &TaobaoVasServiceValidateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVasServiceValidateRequest) GetApiMethodName() string {
    return "taobao.vas.service.validate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVasServiceValidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServCode Setter
// 服务编码
func (r *TaobaoVasServiceValidateRequest) SetServCode(servCode string) error {
    r.servCode = servCode
    r.Set("serv_code", servCode)
    return nil
}

// ServCode Getter
func (r TaobaoVasServiceValidateRequest) GetServCode() string {
    return r.servCode
}
// Nick Setter
// 用户昵称
func (r *TaobaoVasServiceValidateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoVasServiceValidateRequest) GetNick() string {
    return r.nick
}

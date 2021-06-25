package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增值服务订购服务验证 APIRequest
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

func NewTaobaoVasServiceValidateRequest() *TaobaoVasServiceValidateRequest{
    return &TaobaoVasServiceValidateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVasServiceValidateRequest) GetApiMethodName() string {
    return "taobao.vas.service.validate"
}

func (r TaobaoVasServiceValidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoVasServiceValidateRequest) SetServCode(servCode string) error {
    r.servCode = servCode
    r.Set("serv_code", servCode)
    return nil
}

func (r TaobaoVasServiceValidateRequest) GetServCode() string {
    return r.servCode
}

func (r *TaobaoVasServiceValidateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoVasServiceValidateRequest) GetNick() string {
    return r.nick
}


package yunosminiapp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
车载小程序外部服务调用 APIRequest
yunos.miniapp.datatunnel.call

对客户提供的api进行统一封装调用。
*/
type YunosMiniappDatatunnelCallRequest struct {
    model.Params

    // 参数
    param   *BaseRequest 

}

func NewYunosMiniappDatatunnelCallRequest() *YunosMiniappDatatunnelCallRequest{
    return &YunosMiniappDatatunnelCallRequest{
        Params: model.NewParams(),
    }
}

func (r YunosMiniappDatatunnelCallRequest) GetApiMethodName() string {
    return "yunos.miniapp.datatunnel.call"
}

func (r YunosMiniappDatatunnelCallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosMiniappDatatunnelCallRequest) SetParam(param *BaseRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r YunosMiniappDatatunnelCallRequest) GetParam() *BaseRequest {
    return r.param
}


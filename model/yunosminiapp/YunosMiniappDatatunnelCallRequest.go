package yunosminiapp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
车载小程序外部服务调用 API请求
yunos.miniapp.datatunnel.call

对客户提供的api进行统一封装调用。
*/
type YunosMiniappDatatunnelCallRequest struct {
    model.Params
    // 参数
    param   *BaseRequest
}

// 初始化YunosMiniappDatatunnelCallRequest对象
func NewYunosMiniappDatatunnelCallRequest() *YunosMiniappDatatunnelCallRequest{
    return &YunosMiniappDatatunnelCallRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosMiniappDatatunnelCallRequest) GetApiMethodName() string {
    return "yunos.miniapp.datatunnel.call"
}

// IRequest interface 方法, 获取API参数
func (r YunosMiniappDatatunnelCallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 参数
func (r *YunosMiniappDatatunnelCallRequest) SetParam(param *BaseRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r YunosMiniappDatatunnelCallRequest) GetParam() *BaseRequest {
    return r.param
}

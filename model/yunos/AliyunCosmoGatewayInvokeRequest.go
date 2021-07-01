package yunos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
alios cosmo服务调用 API请求
aliyun.cosmo.gateway.invoke

AliOS cosmo服务分发平台对外调用接口
*/
type AliyunCosmoGatewayInvokeAPIRequest struct {
    model.Params
    // 请求上下文参数
    _context   *RdamContext
    // 请求对象
    _rdamRequest   *RdamGenericRequest
}

// 初始化AliyunCosmoGatewayInvokeAPIRequest对象
func NewAliyunCosmoGatewayInvokeRequest() *AliyunCosmoGatewayInvokeAPIRequest{
    return &AliyunCosmoGatewayInvokeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunCosmoGatewayInvokeAPIRequest) GetApiMethodName() string {
    return "aliyun.cosmo.gateway.invoke"
}

// IRequest interface 方法, 获取API参数
func (r AliyunCosmoGatewayInvokeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Context Setter
// 请求上下文参数
func (r *AliyunCosmoGatewayInvokeAPIRequest) SetContext(_context *RdamContext) error {
    r._context = _context
    r.Set("context", _context)
    return nil
}

// Context Getter
func (r AliyunCosmoGatewayInvokeAPIRequest) GetContext() *RdamContext {
    return r._context
}
// RdamRequest Setter
// 请求对象
func (r *AliyunCosmoGatewayInvokeAPIRequest) SetRdamRequest(_rdamRequest *RdamGenericRequest) error {
    r._rdamRequest = _rdamRequest
    r.Set("rdam_request", _rdamRequest)
    return nil
}

// RdamRequest Getter
func (r AliyunCosmoGatewayInvokeAPIRequest) GetRdamRequest() *RdamGenericRequest {
    return r._rdamRequest
}

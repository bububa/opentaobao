package aliospay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取支付token API请求
aliyun.alios.pay.token.get

商户用来获取支付的授权token
*/
type AliyunAliosPayTokenGetAPIRequest struct {
    model.Params
    // 请求参数
    _getTokenRequest   *GetTokenRequest
}

// 初始化AliyunAliosPayTokenGetAPIRequest对象
func NewAliyunAliosPayTokenGetRequest() *AliyunAliosPayTokenGetAPIRequest{
    return &AliyunAliosPayTokenGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunAliosPayTokenGetAPIRequest) GetApiMethodName() string {
    return "aliyun.alios.pay.token.get"
}

// IRequest interface 方法, 获取API参数
func (r AliyunAliosPayTokenGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GetTokenRequest Setter
// 请求参数
func (r *AliyunAliosPayTokenGetAPIRequest) SetGetTokenRequest(_getTokenRequest *GetTokenRequest) error {
    r._getTokenRequest = _getTokenRequest
    r.Set("get_token_request", _getTokenRequest)
    return nil
}

// GetTokenRequest Getter
func (r AliyunAliosPayTokenGetAPIRequest) GetGetTokenRequest() *GetTokenRequest {
    return r._getTokenRequest
}

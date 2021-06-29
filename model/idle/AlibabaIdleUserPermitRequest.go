package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户appkey授权 API请求
alibaba.idle.user.permit

用于记录登录用户与服务商的绑定关系，用于业务数据分发和授权校验
*/
type AlibabaIdleUserPermitRequest struct {
    model.Params
    // 授权请求
    _paramUserGrantRequest   *UserGrantRequest
}

// 初始化AlibabaIdleUserPermitRequest对象
func NewAlibabaIdleUserPermitRequest() *AlibabaIdleUserPermitRequest{
    return &AlibabaIdleUserPermitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleUserPermitRequest) GetApiMethodName() string {
    return "alibaba.idle.user.permit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleUserPermitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamUserGrantRequest Setter
// 授权请求
func (r *AlibabaIdleUserPermitRequest) SetParamUserGrantRequest(_paramUserGrantRequest *UserGrantRequest) error {
    r._paramUserGrantRequest = _paramUserGrantRequest
    r.Set("param_user_grant_request", _paramUserGrantRequest)
    return nil
}

// ParamUserGrantRequest Getter
func (r AlibabaIdleUserPermitRequest) GetParamUserGrantRequest() *UserGrantRequest {
    return r._paramUserGrantRequest
}

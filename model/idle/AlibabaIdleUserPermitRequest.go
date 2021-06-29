package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户appkey授权 APIRequest
alibaba.idle.user.permit

用于记录登录用户与服务商的绑定关系，用于业务数据分发和授权校验
*/
type AlibabaIdleUserPermitRequest struct {
    model.Params

    // 授权请求
    paramUserGrantRequest   *UserGrantRequest 

}

func NewAlibabaIdleUserPermitRequest() *AlibabaIdleUserPermitRequest{
    return &AlibabaIdleUserPermitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleUserPermitRequest) GetApiMethodName() string {
    return "alibaba.idle.user.permit"
}

func (r AlibabaIdleUserPermitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleUserPermitRequest) SetParamUserGrantRequest(paramUserGrantRequest *UserGrantRequest) error {
    r.paramUserGrantRequest = paramUserGrantRequest
    r.Set("param_user_grant_request", paramUserGrantRequest)
    return nil
}

func (r AlibabaIdleUserPermitRequest) GetParamUserGrantRequest() *UserGrantRequest {
    return r.paramUserGrantRequest
}


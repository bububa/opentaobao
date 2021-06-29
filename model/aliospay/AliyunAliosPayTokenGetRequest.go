package aliospay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取支付token APIRequest
aliyun.alios.pay.token.get

商户用来获取支付的授权token
*/
type AliyunAliosPayTokenGetRequest struct {
    model.Params

    // 请求参数
    getTokenRequest   *GetTokenRequest 

}

func NewAliyunAliosPayTokenGetRequest() *AliyunAliosPayTokenGetRequest{
    return &AliyunAliosPayTokenGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliyunAliosPayTokenGetRequest) GetApiMethodName() string {
    return "aliyun.alios.pay.token.get"
}

func (r AliyunAliosPayTokenGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliyunAliosPayTokenGetRequest) SetGetTokenRequest(getTokenRequest *GetTokenRequest) error {
    r.getTokenRequest = getTokenRequest
    r.Set("get_token_request", getTokenRequest)
    return nil
}

func (r AliyunAliosPayTokenGetRequest) GetGetTokenRequest() *GetTokenRequest {
    return r.getTokenRequest
}


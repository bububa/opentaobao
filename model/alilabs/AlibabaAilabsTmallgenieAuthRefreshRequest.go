package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
刷新token APIRequest
alibaba.ailabs.tmallgenie.auth.refresh

通过此接口刷新天猫精灵授权token
*/
type AlibabaAilabsTmallgenieAuthRefreshRequest struct {
    model.Params

    // refresh_token_request
    refreshTokenRequest   *TopRefreshReqDto 

}

func NewAlibabaAilabsTmallgenieAuthRefreshRequest() *AlibabaAilabsTmallgenieAuthRefreshRequest{
    return &AlibabaAilabsTmallgenieAuthRefreshRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsTmallgenieAuthRefreshRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.refresh"
}

func (r AlibabaAilabsTmallgenieAuthRefreshRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsTmallgenieAuthRefreshRequest) SetRefreshTokenRequest(refreshTokenRequest *TopRefreshReqDto) error {
    r.refreshTokenRequest = refreshTokenRequest
    r.Set("refresh_token_request", refreshTokenRequest)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthRefreshRequest) GetRefreshTokenRequest() *TopRefreshReqDto {
    return r.refreshTokenRequest
}


package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
刷新token API请求
alibaba.ailabs.tmallgenie.auth.refresh

通过此接口刷新天猫精灵授权token
*/
type AlibabaAilabsTmallgenieAuthRefreshAPIRequest struct {
    model.Params
    // refresh_token_request
    _refreshTokenRequest   *TopRefreshReqDTO
}

// 初始化AlibabaAilabsTmallgenieAuthRefreshAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthRefreshRequest() *AlibabaAilabsTmallgenieAuthRefreshAPIRequest{
    return &AlibabaAilabsTmallgenieAuthRefreshAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthRefreshAPIRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.refresh"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthRefreshAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefreshTokenRequest Setter
// refresh_token_request
func (r *AlibabaAilabsTmallgenieAuthRefreshAPIRequest) SetRefreshTokenRequest(_refreshTokenRequest *TopRefreshReqDTO) error {
    r._refreshTokenRequest = _refreshTokenRequest
    r.Set("refresh_token_request", _refreshTokenRequest)
    return nil
}

// RefreshTokenRequest Getter
func (r AlibabaAilabsTmallgenieAuthRefreshAPIRequest) GetRefreshTokenRequest() *TopRefreshReqDTO {
    return r._refreshTokenRequest
}

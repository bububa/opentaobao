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
type AlibabaAilabsTmallgenieAuthRefreshRequest struct {
    model.Params
    // refresh_token_request
    _refreshTokenRequest   *TopRefreshReqDto
}

// 初始化AlibabaAilabsTmallgenieAuthRefreshRequest对象
func NewAlibabaAilabsTmallgenieAuthRefreshRequest() *AlibabaAilabsTmallgenieAuthRefreshRequest{
    return &AlibabaAilabsTmallgenieAuthRefreshRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthRefreshRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.refresh"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthRefreshRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefreshTokenRequest Setter
// refresh_token_request
func (r *AlibabaAilabsTmallgenieAuthRefreshRequest) SetRefreshTokenRequest(_refreshTokenRequest *TopRefreshReqDto) error {
    r._refreshTokenRequest = _refreshTokenRequest
    r.Set("refresh_token_request", _refreshTokenRequest)
    return nil
}

// RefreshTokenRequest Getter
func (r AlibabaAilabsTmallgenieAuthRefreshRequest) GetRefreshTokenRequest() *TopRefreshReqDto {
    return r._refreshTokenRequest
}

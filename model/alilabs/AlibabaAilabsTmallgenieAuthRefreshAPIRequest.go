package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthRefreshAPIRequest 刷新token API请求
// alibaba.ailabs.tmallgenie.auth.refresh
//
// 通过此接口刷新天猫精灵授权token
type AlibabaAilabsTmallgenieAuthRefreshAPIRequest struct {
	model.Params
	// refresh_token_request
	_refreshTokenRequest *TopRefreshReqDto
}

// NewAlibabaAilabsTmallgenieAuthRefreshRequest 初始化AlibabaAilabsTmallgenieAuthRefreshAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthRefreshRequest() *AlibabaAilabsTmallgenieAuthRefreshAPIRequest {
	return &AlibabaAilabsTmallgenieAuthRefreshAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthRefreshAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.refresh"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthRefreshAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsTmallgenieAuthRefreshAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefreshTokenRequest is RefreshTokenRequest Setter
// refresh_token_request
func (r *AlibabaAilabsTmallgenieAuthRefreshAPIRequest) SetRefreshTokenRequest(_refreshTokenRequest *TopRefreshReqDto) error {
	r._refreshTokenRequest = _refreshTokenRequest
	r.Set("refresh_token_request", _refreshTokenRequest)
	return nil
}

// GetRefreshTokenRequest RefreshTokenRequest Getter
func (r AlibabaAilabsTmallgenieAuthRefreshAPIRequest) GetRefreshTokenRequest() *TopRefreshReqDto {
	return r._refreshTokenRequest
}

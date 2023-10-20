package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugdownloadgetentauthentAPIRequest 获取授权企业列表 API请求
// alibaba.alihealth.drug.download.getentauthent
//
// D2D数据落地获取授权企业列表
type AlibabaalihealthdrugdownloadgetentauthentAPIRequest struct {
	model.Params
	// 授权开始时间
	_authBeginDate string
	// 授权结束时间
	_authEndDate string
}

// NewAlibabaalihealthdrugdownloadgetentauthentRequest 初始化AlibabaalihealthdrugdownloadgetentauthentAPIRequest对象
func NewAlibabaalihealthdrugdownloadgetentauthentRequest() *AlibabaalihealthdrugdownloadgetentauthentAPIRequest {
	return &AlibabaalihealthdrugdownloadgetentauthentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugdownloadgetentauthentAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.download.getentauthent"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugdownloadgetentauthentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugdownloadgetentauthentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuthBeginDate is AuthBeginDate Setter
// 授权开始时间
func (r *AlibabaalihealthdrugdownloadgetentauthentAPIRequest) SetAuthBeginDate(_authBeginDate string) error {
	r._authBeginDate = _authBeginDate
	r.Set("auth_begin_date", _authBeginDate)
	return nil
}

// GetAuthBeginDate AuthBeginDate Getter
func (r AlibabaalihealthdrugdownloadgetentauthentAPIRequest) GetAuthBeginDate() string {
	return r._authBeginDate
}

// SetAuthEndDate is AuthEndDate Setter
// 授权结束时间
func (r *AlibabaalihealthdrugdownloadgetentauthentAPIRequest) SetAuthEndDate(_authEndDate string) error {
	r._authEndDate = _authEndDate
	r.Set("auth_end_date", _authEndDate)
	return nil
}

// GetAuthEndDate AuthEndDate Getter
func (r AlibabaalihealthdrugdownloadgetentauthentAPIRequest) GetAuthEndDate() string {
	return r._authEndDate
}

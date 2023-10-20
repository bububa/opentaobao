package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugdownloadentlistAPIRequest 企业下载列表 API请求
// alibaba.alihealth.drug.download.entlist
//
// 获取企业的下载文件列表
type AlibabaalihealthdrugdownloadentlistAPIRequest struct {
	model.Params
	// appKey
	_appKeyN string
}

// NewAlibabaalihealthdrugdownloadentlistRequest 初始化AlibabaalihealthdrugdownloadentlistAPIRequest对象
func NewAlibabaalihealthdrugdownloadentlistRequest() *AlibabaalihealthdrugdownloadentlistAPIRequest {
	return &AlibabaalihealthdrugdownloadentlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugdownloadentlistAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.download.entlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugdownloadentlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugdownloadentlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppKeyN is AppKeyN Setter
// appKey
func (r *AlibabaalihealthdrugdownloadentlistAPIRequest) SetAppKeyN(_appKeyN string) error {
	r._appKeyN = _appKeyN
	r.Set("app_key_n", _appKeyN)
	return nil
}

// GetAppKeyN AppKeyN Getter
func (r AlibabaalihealthdrugdownloadentlistAPIRequest) GetAppKeyN() string {
	return r._appKeyN
}

package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunososupdateappversionupdateAPIRequest 应用升级任务更新 API请求
// yunos.osupdate.appversion.update
//
// 应用升级任务更新
type YunososupdateappversionupdateAPIRequest struct {
	model.Params
	// 应用版本升级信息
	_appVersion *TvAppVersion
}

// NewYunososupdateappversionupdateRequest 初始化YunososupdateappversionupdateAPIRequest对象
func NewYunososupdateappversionupdateRequest() *YunososupdateappversionupdateAPIRequest {
	return &YunososupdateappversionupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunososupdateappversionupdateAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.appversion.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunososupdateappversionupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunososupdateappversionupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppVersion is AppVersion Setter
// 应用版本升级信息
func (r *YunososupdateappversionupdateAPIRequest) SetAppVersion(_appVersion *TvAppVersion) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// GetAppVersion AppVersion Getter
func (r YunososupdateappversionupdateAPIRequest) GetAppVersion() *TvAppVersion {
	return r._appVersion
}

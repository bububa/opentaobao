package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunososupdateappversioncreateAPIRequest 创建应用升级任务 API请求
// yunos.osupdate.appversion.create
//
// 创建应用升级任务
type YunososupdateappversioncreateAPIRequest struct {
	model.Params
	// 应用版本信息
	_appVersion *TvAppVersion
}

// NewYunososupdateappversioncreateRequest 初始化YunososupdateappversioncreateAPIRequest对象
func NewYunososupdateappversioncreateRequest() *YunososupdateappversioncreateAPIRequest {
	return &YunososupdateappversioncreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunososupdateappversioncreateAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.appversion.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunososupdateappversioncreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunososupdateappversioncreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppVersion is AppVersion Setter
// 应用版本信息
func (r *YunososupdateappversioncreateAPIRequest) SetAppVersion(_appVersion *TvAppVersion) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// GetAppVersion AppVersion Getter
func (r YunososupdateappversioncreateAPIRequest) GetAppVersion() *TvAppVersion {
	return r._appVersion
}

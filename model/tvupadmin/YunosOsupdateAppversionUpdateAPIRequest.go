package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosOsupdateAppversionUpdateAPIRequest
应用升级任务更新 API请求
yunos.osupdate.appversion.update

应用升级任务更新 */
type YunosOsupdateAppversionUpdateAPIRequest struct {
	model.Params
	// 应用版本升级信息
	_appVersion *TvAppVersion
}

// NewYunosOsupdateAppversionUpdateRequest 初始化YunosOsupdateAppversionUpdateAPIRequest对象
func NewYunosOsupdateAppversionUpdateRequest() *YunosOsupdateAppversionUpdateAPIRequest {
	return &YunosOsupdateAppversionUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosOsupdateAppversionUpdateAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.appversion.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosOsupdateAppversionUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AppVersion Setter
// 应用版本升级信息
func (r *YunosOsupdateAppversionUpdateAPIRequest) SetAppVersion(_appVersion *TvAppVersion) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// Get AppVersion Getter
func (r YunosOsupdateAppversionUpdateAPIRequest) GetAppVersion() *TvAppVersion {
	return r._appVersion
}

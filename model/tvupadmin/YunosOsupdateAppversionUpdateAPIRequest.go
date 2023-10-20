package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateAppversionUpdateAPIRequest 应用升级任务更新 API请求
// yunos.osupdate.appversion.update
//
// 应用升级任务更新
type YunosOsupdateAppversionUpdateAPIRequest struct {
	model.Params
	// 应用版本升级信息
	_appVersion *TvAppVersion
}

// NewYunosOsupdateAppversionUpdateRequest 初始化YunosOsupdateAppversionUpdateAPIRequest对象
func NewYunosOsupdateAppversionUpdateRequest() *YunosOsupdateAppversionUpdateAPIRequest {
	return &YunosOsupdateAppversionUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosOsupdateAppversionUpdateAPIRequest) Reset() {
	r._appVersion = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosOsupdateAppversionUpdateAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.appversion.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosOsupdateAppversionUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosOsupdateAppversionUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppVersion is AppVersion Setter
// 应用版本升级信息
func (r *YunosOsupdateAppversionUpdateAPIRequest) SetAppVersion(_appVersion *TvAppVersion) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// GetAppVersion AppVersion Getter
func (r YunosOsupdateAppversionUpdateAPIRequest) GetAppVersion() *TvAppVersion {
	return r._appVersion
}

var poolYunosOsupdateAppversionUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosOsupdateAppversionUpdateRequest()
	},
}

// GetYunosOsupdateAppversionUpdateRequest 从 sync.Pool 获取 YunosOsupdateAppversionUpdateAPIRequest
func GetYunosOsupdateAppversionUpdateAPIRequest() *YunosOsupdateAppversionUpdateAPIRequest {
	return poolYunosOsupdateAppversionUpdateAPIRequest.Get().(*YunosOsupdateAppversionUpdateAPIRequest)
}

// ReleaseYunosOsupdateAppversionUpdateAPIRequest 将 YunosOsupdateAppversionUpdateAPIRequest 放入 sync.Pool
func ReleaseYunosOsupdateAppversionUpdateAPIRequest(v *YunosOsupdateAppversionUpdateAPIRequest) {
	v.Reset()
	poolYunosOsupdateAppversionUpdateAPIRequest.Put(v)
}

package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateAppversionCreateAPIRequest 创建应用升级任务 API请求
// yunos.osupdate.appversion.create
//
// 创建应用升级任务
type YunosOsupdateAppversionCreateAPIRequest struct {
	model.Params
	// 应用版本信息
	_appVersion *TvAppVersion
}

// NewYunosOsupdateAppversionCreateRequest 初始化YunosOsupdateAppversionCreateAPIRequest对象
func NewYunosOsupdateAppversionCreateRequest() *YunosOsupdateAppversionCreateAPIRequest {
	return &YunosOsupdateAppversionCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosOsupdateAppversionCreateAPIRequest) Reset() {
	r._appVersion = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosOsupdateAppversionCreateAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.appversion.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosOsupdateAppversionCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosOsupdateAppversionCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppVersion is AppVersion Setter
// 应用版本信息
func (r *YunosOsupdateAppversionCreateAPIRequest) SetAppVersion(_appVersion *TvAppVersion) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// GetAppVersion AppVersion Getter
func (r YunosOsupdateAppversionCreateAPIRequest) GetAppVersion() *TvAppVersion {
	return r._appVersion
}

var poolYunosOsupdateAppversionCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosOsupdateAppversionCreateRequest()
	},
}

// GetYunosOsupdateAppversionCreateRequest 从 sync.Pool 获取 YunosOsupdateAppversionCreateAPIRequest
func GetYunosOsupdateAppversionCreateAPIRequest() *YunosOsupdateAppversionCreateAPIRequest {
	return poolYunosOsupdateAppversionCreateAPIRequest.Get().(*YunosOsupdateAppversionCreateAPIRequest)
}

// ReleaseYunosOsupdateAppversionCreateAPIRequest 将 YunosOsupdateAppversionCreateAPIRequest 放入 sync.Pool
func ReleaseYunosOsupdateAppversionCreateAPIRequest(v *YunosOsupdateAppversionCreateAPIRequest) {
	v.Reset()
	poolYunosOsupdateAppversionCreateAPIRequest.Put(v)
}

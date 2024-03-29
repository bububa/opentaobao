package yunosappstore

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosAppstoreAppsGetAPIRequest 根据包名列表获取应用信息列表 API请求
// yunos.appstore.apps.get
//
// 根据包名列表获取应用信息列表
type YunosAppstoreAppsGetAPIRequest struct {
	model.Params
	// 应用包名列表
	_pkgs []string
}

// NewYunosAppstoreAppsGetRequest 初始化YunosAppstoreAppsGetAPIRequest对象
func NewYunosAppstoreAppsGetRequest() *YunosAppstoreAppsGetAPIRequest {
	return &YunosAppstoreAppsGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosAppstoreAppsGetAPIRequest) Reset() {
	r._pkgs = r._pkgs[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosAppstoreAppsGetAPIRequest) GetApiMethodName() string {
	return "yunos.appstore.apps.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosAppstoreAppsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosAppstoreAppsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPkgs is Pkgs Setter
// 应用包名列表
func (r *YunosAppstoreAppsGetAPIRequest) SetPkgs(_pkgs []string) error {
	r._pkgs = _pkgs
	r.Set("pkgs", _pkgs)
	return nil
}

// GetPkgs Pkgs Getter
func (r YunosAppstoreAppsGetAPIRequest) GetPkgs() []string {
	return r._pkgs
}

var poolYunosAppstoreAppsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosAppstoreAppsGetRequest()
	},
}

// GetYunosAppstoreAppsGetRequest 从 sync.Pool 获取 YunosAppstoreAppsGetAPIRequest
func GetYunosAppstoreAppsGetAPIRequest() *YunosAppstoreAppsGetAPIRequest {
	return poolYunosAppstoreAppsGetAPIRequest.Get().(*YunosAppstoreAppsGetAPIRequest)
}

// ReleaseYunosAppstoreAppsGetAPIRequest 将 YunosAppstoreAppsGetAPIRequest 放入 sync.Pool
func ReleaseYunosAppstoreAppsGetAPIRequest(v *YunosAppstoreAppsGetAPIRequest) {
	v.Reset()
	poolYunosAppstoreAppsGetAPIRequest.Put(v)
}

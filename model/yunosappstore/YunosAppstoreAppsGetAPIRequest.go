package yunosappstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosappstoreappsgetAPIRequest 根据包名列表获取应用信息列表 API请求
// yunos.appstore.apps.get
//
// 根据包名列表获取应用信息列表
type YunosappstoreappsgetAPIRequest struct {
	model.Params
	// 应用包名列表
	_pkgs []string
}

// NewYunosappstoreappsgetRequest 初始化YunosappstoreappsgetAPIRequest对象
func NewYunosappstoreappsgetRequest() *YunosappstoreappsgetAPIRequest {
	return &YunosappstoreappsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosappstoreappsgetAPIRequest) GetApiMethodName() string {
	return "yunos.appstore.apps.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosappstoreappsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosappstoreappsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPkgs is Pkgs Setter
// 应用包名列表
func (r *YunosappstoreappsgetAPIRequest) SetPkgs(_pkgs []string) error {
	r._pkgs = _pkgs
	r.Set("pkgs", _pkgs)
	return nil
}

// GetPkgs Pkgs Getter
func (r YunosappstoreappsgetAPIRequest) GetPkgs() []string {
	return r._pkgs
}

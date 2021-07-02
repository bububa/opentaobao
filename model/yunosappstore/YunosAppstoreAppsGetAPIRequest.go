package yunosappstore

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosAppstoreAppsGetAPIRequest) GetApiMethodName() string {
	return "yunos.appstore.apps.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosAppstoreAppsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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

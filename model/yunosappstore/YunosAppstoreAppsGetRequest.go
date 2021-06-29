package yunosappstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据包名列表获取应用信息列表 API请求
yunos.appstore.apps.get

根据包名列表获取应用信息列表
*/
type YunosAppstoreAppsGetRequest struct {
    model.Params
    // 应用包名列表
    _pkgs   []string
}

// 初始化YunosAppstoreAppsGetRequest对象
func NewYunosAppstoreAppsGetRequest() *YunosAppstoreAppsGetRequest{
    return &YunosAppstoreAppsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosAppstoreAppsGetRequest) GetApiMethodName() string {
    return "yunos.appstore.apps.get"
}

// IRequest interface 方法, 获取API参数
func (r YunosAppstoreAppsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Pkgs Setter
// 应用包名列表
func (r *YunosAppstoreAppsGetRequest) SetPkgs(_pkgs []string) error {
    r._pkgs = _pkgs
    r.Set("pkgs", _pkgs)
    return nil
}

// Pkgs Getter
func (r YunosAppstoreAppsGetRequest) GetPkgs() []string {
    return r._pkgs
}
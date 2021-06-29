package yunosappstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据包名列表获取应用信息列表 APIRequest
yunos.appstore.apps.get

根据包名列表获取应用信息列表
*/
type YunosAppstoreAppsGetRequest struct {
    model.Params

    // 应用包名列表
    pkgs   []string 

}

func NewYunosAppstoreAppsGetRequest() *YunosAppstoreAppsGetRequest{
    return &YunosAppstoreAppsGetRequest{
        Params: model.NewParams(),
    }
}

func (r YunosAppstoreAppsGetRequest) GetApiMethodName() string {
    return "yunos.appstore.apps.get"
}

func (r YunosAppstoreAppsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosAppstoreAppsGetRequest) SetPkgs(pkgs []string) error {
    r.pkgs = pkgs
    r.Set("pkgs", pkgs)
    return nil
}

func (r YunosAppstoreAppsGetRequest) GetPkgs() []string {
    return r.pkgs
}


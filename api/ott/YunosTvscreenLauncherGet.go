package ott

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ott"
)

/* 
一体机桌面 
yunos.tvscreen.launcher.get

LCTS一体机桌面后台,提供基于运营坑位适配的桌面服务
*/
func YunosTvscreenLauncherGet(clt *core.SDKClient, req *ott.YunosTvscreenLauncherGetRequest, session string) (*ott.YunosTvscreenLauncherGetAPIResponse, error) {
    var resp ott.YunosTvscreenLauncherGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

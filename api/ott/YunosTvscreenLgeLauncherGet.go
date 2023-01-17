package ott

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ott"
)

// YunosTvscreenLgeLauncherGet LG用桌面信息获取
// yunos.tvscreen.lge.launcher.get
//
// LCTS一体机桌面后台,提供基于运营坑位适配的桌面服务,根据LG的需求定制输出格式
func YunosTvscreenLgeLauncherGet(clt *core.SDKClient, req *ott.YunosTvscreenLgeLauncherGetAPIRequest, session string) (*ott.YunosTvscreenLgeLauncherGetAPIResponse, error) {
	var resp ott.YunosTvscreenLgeLauncherGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

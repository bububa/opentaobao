package ott

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ott"
)

// Yunostvscreenlauncherget 一体机桌面
// yunos.tvscreen.launcher.get
//
// LCTS一体机桌面后台,提供基于运营坑位适配的桌面服务
func Yunostvscreenlauncherget(clt *core.SDKClient, req *ott.YunostvscreenlaunchergetAPIRequest, session string) (*ott.YunostvscreenlaunchergetAPIResponse, error) {
	var resp ott.YunostvscreenlaunchergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

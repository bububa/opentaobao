package yunosappstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosappstore"
)

// Yunosappstoreappsget 根据包名列表获取应用信息列表
// yunos.appstore.apps.get
//
// 根据包名列表获取应用信息列表
func Yunosappstoreappsget(clt *core.SDKClient, req *yunosappstore.YunosappstoreappsgetAPIRequest, session string) (*yunosappstore.YunosappstoreappsgetAPIResponse, error) {
	var resp yunosappstore.YunosappstoreappsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

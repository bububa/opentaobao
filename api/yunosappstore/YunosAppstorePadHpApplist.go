package yunosappstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosappstore"
)

// Yunosappstorepadhpapplist 查询HpPad appList
// yunos.appstore.pad.hp.applist
//
// 提供hp pad应用群数据
func Yunosappstorepadhpapplist(clt *core.SDKClient, req *yunosappstore.YunosappstorepadhpapplistAPIRequest, session string) (*yunosappstore.YunosappstorepadhpapplistAPIResponse, error) {
	var resp yunosappstore.YunosappstorepadhpapplistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

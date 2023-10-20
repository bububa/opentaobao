package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// Taobaominiappextdeliveryappchannelsync ISV写入应用的渠道信息
// taobao.miniapp.ext.delivery.app.channel.sync
//
// ISV写入应用的渠道信息
func Taobaominiappextdeliveryappchannelsync(clt *core.SDKClient, req *miniapp.TaobaominiappextdeliveryappchannelsyncAPIRequest, session string) (*miniapp.TaobaominiappextdeliveryappchannelsyncAPIResponse, error) {
	var resp miniapp.TaobaominiappextdeliveryappchannelsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

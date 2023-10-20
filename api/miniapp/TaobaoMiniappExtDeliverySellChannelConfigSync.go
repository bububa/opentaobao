package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// Taobaominiappextdeliverysellchannelconfigsync 写入商家配置信息
// taobao.miniapp.ext.delivery.sell.channel.config.sync
//
// 写入商家配置信息
func Taobaominiappextdeliverysellchannelconfigsync(clt *core.SDKClient, req *miniapp.TaobaominiappextdeliverysellchannelconfigsyncAPIRequest, session string) (*miniapp.TaobaominiappextdeliverysellchannelconfigsyncAPIResponse, error) {
	var resp miniapp.TaobaominiappextdeliverysellchannelconfigsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

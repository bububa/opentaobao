package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// Taobaominiappextdeliveryappchannelconfigsquery ISV查询应用的渠道信息
// taobao.miniapp.ext.delivery.app.channel.configs.query
//
// ISV查询应用的渠道信息
func Taobaominiappextdeliveryappchannelconfigsquery(clt *core.SDKClient, req *miniapp.TaobaominiappextdeliveryappchannelconfigsqueryAPIRequest, session string) (*miniapp.TaobaominiappextdeliveryappchannelconfigsqueryAPIResponse, error) {
	var resp miniapp.TaobaominiappextdeliveryappchannelconfigsqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

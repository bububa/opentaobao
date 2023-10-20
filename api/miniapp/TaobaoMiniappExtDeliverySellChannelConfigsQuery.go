package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// Taobaominiappextdeliverysellchannelconfigsquery 查询商家配置的信息
// taobao.miniapp.ext.delivery.sell.channel.configs.query
//
// 查询商家配置的信息
func Taobaominiappextdeliverysellchannelconfigsquery(clt *core.SDKClient, req *miniapp.TaobaominiappextdeliverysellchannelconfigsqueryAPIRequest, session string) (*miniapp.TaobaominiappextdeliverysellchannelconfigsqueryAPIResponse, error) {
	var resp miniapp.TaobaominiappextdeliverysellchannelconfigsqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

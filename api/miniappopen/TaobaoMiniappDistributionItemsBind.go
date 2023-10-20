package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// Taobaominiappdistributionitemsbind 【已废弃】小程序投放-商品绑定/解绑
// taobao.miniapp.distribution.items.bind
//
// 【已废弃，请使用 taobao.miniapp.distribution.order.items.bind 接口】提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。
func Taobaominiappdistributionitemsbind(clt *core.SDKClient, req *miniappopen.TaobaominiappdistributionitemsbindAPIRequest, session string) (*miniappopen.TaobaominiappdistributionitemsbindAPIResponse, error) {
	var resp miniappopen.TaobaominiappdistributionitemsbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

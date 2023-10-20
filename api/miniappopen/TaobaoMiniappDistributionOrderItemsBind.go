package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// Taobaominiappdistributionorderitemsbind 小程序投放-基于投放计划绑定/解绑商品
// taobao.miniapp.distribution.order.items.bind
//
// 提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。
func Taobaominiappdistributionorderitemsbind(clt *core.SDKClient, req *miniappopen.TaobaominiappdistributionorderitemsbindAPIRequest, session string) (*miniappopen.TaobaominiappdistributionorderitemsbindAPIResponse, error) {
	var resp miniappopen.TaobaominiappdistributionorderitemsbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappDistributionItemsBind 小程序投放-商品绑定/解绑
// taobao.miniapp.distribution.items.bind
//
// 提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。
func TaobaoMiniappDistributionItemsBind(clt *core.SDKClient, req *miniappopen.TaobaoMiniappDistributionItemsBindAPIRequest, session string) (*miniappopen.TaobaoMiniappDistributionItemsBindAPIResponse, error) {
	var resp miniappopen.TaobaoMiniappDistributionItemsBindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

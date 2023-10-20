package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappVirtualItemGet 小程序关联虚拟商品查询
// taobao.miniapp.virtual.item.get
//
// 小程序关联虚拟商品查询
func TaobaoMiniappVirtualItemGet(clt *core.SDKClient, req *miniappopen.TaobaoMiniappVirtualItemGetAPIRequest, session string) (*miniappopen.TaobaoMiniappVirtualItemGetAPIResponse, error) {
	var resp miniappopen.TaobaoMiniappVirtualItemGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappDistributionOrderPrecreate 代商家预创建投放计划
// taobao.miniapp.distribution.order.precreate
//
// 帮助商家，预创建小程序的投放计划，预创建的投放计划，在商家确认以后，则会生效可用。
func TaobaoMiniappDistributionOrderPrecreate(clt *core.SDKClient, req *miniappopen.TaobaoMiniappDistributionOrderPrecreateAPIRequest, session string) (*miniappopen.TaobaoMiniappDistributionOrderPrecreateAPIResponse, error) {
	var resp miniappopen.TaobaoMiniappDistributionOrderPrecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

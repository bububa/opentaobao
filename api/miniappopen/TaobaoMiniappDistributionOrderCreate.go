package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappDistributionOrderCreate 创建小程序投放计划
// taobao.miniapp.distribution.order.create
//
// 帮助商家，创建小程序的投放计划。该api，仅针对特定场景开放，目前仅支持客服场景，具体支持的场景列表请联系技术支持或业务对接人确认。
func TaobaoMiniappDistributionOrderCreate(clt *core.SDKClient, req *miniappopen.TaobaoMiniappDistributionOrderCreateAPIRequest, session string) (*miniappopen.TaobaoMiniappDistributionOrderCreateAPIResponse, error) {
	var resp miniappopen.TaobaoMiniappDistributionOrderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

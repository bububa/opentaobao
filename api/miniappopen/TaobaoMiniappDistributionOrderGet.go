package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

/* TaobaoMiniappDistributionOrderGet
小程序投放-查询小程序投放计划信息
taobao.miniapp.distribution.order.get

服务商可通过该API，读取自己开发的小程序对应的投放计划的相关信息 */
func TaobaoMiniappDistributionOrderGet(clt *core.SDKClient, req *miniappopen.TaobaoMiniappDistributionOrderGetAPIRequest, session string) (*miniappopen.TaobaoMiniappDistributionOrderGetAPIResponse, error) {
	var resp miniappopen.TaobaoMiniappDistributionOrderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

/* TaobaoWlbWmsInventoryProfitlossGet
通过订单列表批量获取库存损益单信息
taobao.wlb.wms.inventory.profitloss.get

通过订单列表批量获取库存损益单信息 */
func TaobaoWlbWmsInventoryProfitlossGet(clt *core.SDKClient, req *wms.TaobaoWlbWmsInventoryProfitlossGetAPIRequest, session string) (*wms.TaobaoWlbWmsInventoryProfitlossGetAPIResponse, error) {
	var resp wms.TaobaoWlbWmsInventoryProfitlossGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

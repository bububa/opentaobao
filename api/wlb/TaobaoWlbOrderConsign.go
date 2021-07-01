package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

/* TaobaoWlbOrderConsign
物流宝订单已发货通知接口
taobao.wlb.order.consign

如果erp导入淘宝交易订单到物流宝，当物流宝订单已发货的时候，erp需要调用该接口来通知物流订单和淘宝交易订单已发货 */
func TaobaoWlbOrderConsign(clt *core.SDKClient, req *wlb.TaobaoWlbOrderConsignAPIRequest, session string) (*wlb.TaobaoWlbOrderConsignAPIResponse, error) {
	var resp wlb.TaobaoWlbOrderConsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

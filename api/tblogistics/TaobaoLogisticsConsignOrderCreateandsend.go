package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaologisticsconsignordercreateandsend 创建订单并发货
// taobao.logistics.consign.order.createandsend
//
// 创建物流订单，并发货。
func Taobaologisticsconsignordercreateandsend(clt *core.SDKClient, req *tblogistics.TaobaologisticsconsignordercreateandsendAPIRequest, session string) (*tblogistics.TaobaologisticsconsignordercreateandsendAPIResponse, error) {
	var resp tblogistics.TaobaologisticsconsignordercreateandsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

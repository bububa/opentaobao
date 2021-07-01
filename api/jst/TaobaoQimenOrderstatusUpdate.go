package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

/* TaobaoQimenOrderstatusUpdate
订单状态更新接口
taobao.qimen.orderstatus.update

星盘和ISV，可以通过此接口，来更新订单状态。此接口应用于使用阿里星盘分单，且使用商家系统（非阿里掌柜）接单/拒单的模式下更新订单状态。 */
func TaobaoQimenOrderstatusUpdate(clt *core.SDKClient, req *jst.TaobaoQimenOrderstatusUpdateAPIRequest, session string) (*jst.TaobaoQimenOrderstatusUpdateAPIResponse, error) {
	var resp jst.TaobaoQimenOrderstatusUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

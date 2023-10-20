package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaofuwupurchaseorderconfirm 服务市场内购服务下单接口
// taobao.fuwu.purchase.order.confirm
//
// 通过传入服务市场商品的itemcode等信息，返回给服务商内购服务的下单链接
func Taobaofuwupurchaseorderconfirm(clt *core.SDKClient, req *servicecenter.TaobaofuwupurchaseorderconfirmAPIRequest, session string) (*servicecenter.TaobaofuwupurchaseorderconfirmAPIResponse, error) {
	var resp servicecenter.TaobaofuwupurchaseorderconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package tvpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

/* TaobaoTvpayOrderQuery
tv支付查询订单状态
taobao.tvpay.order.query

tv支付查询订单状态 */
func TaobaoTvpayOrderQuery(clt *core.SDKClient, req *tvpay.TaobaoTvpayOrderQueryAPIRequest, session string) (*tvpay.TaobaoTvpayOrderQueryAPIResponse, error) {
	var resp tvpay.TaobaoTvpayOrderQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

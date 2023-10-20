package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// TaobaoAlitripSellerRefundmoneyConfirm 【机票代理商订单】确认退款
// taobao.alitrip.seller.refundmoney.confirm
//
// 代理人确认退票申请单的退款
func TaobaoAlitripSellerRefundmoneyConfirm(clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerRefundmoneyConfirmAPIRequest, resp *jipiao.TaobaoAlitripSellerRefundmoneyConfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

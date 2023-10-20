package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// TaobaoAlitripSellerRefundConfirmreturn 【机票代理商】确认退票
// taobao.alitrip.seller.refund.confirmreturn
//
// 确认退票
func TaobaoAlitripSellerRefundConfirmreturn(clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerRefundConfirmreturnAPIRequest, resp *jipiao.TaobaoAlitripSellerRefundConfirmreturnAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

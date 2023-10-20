package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// TaobaoAlitripSellerRefundRefusereturn 【机票代理商】拒绝退票
// taobao.alitrip.seller.refund.refusereturn
//
// 拒绝退票
func TaobaoAlitripSellerRefundRefusereturn(clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerRefundRefusereturnAPIRequest, resp *jipiao.TaobaoAlitripSellerRefundRefusereturnAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

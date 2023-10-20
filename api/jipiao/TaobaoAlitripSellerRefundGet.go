package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// TaobaoAlitripSellerRefundGet 【机票代理商】退票申请单详情
// taobao.alitrip.seller.refund.get
//
// 查询退票申请单详情
func TaobaoAlitripSellerRefundGet(clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerRefundGetAPIRequest, resp *jipiao.TaobaoAlitripSellerRefundGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

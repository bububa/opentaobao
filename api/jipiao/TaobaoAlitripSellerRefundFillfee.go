package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// TaobaoAlitripSellerRefundFillfee 机票代理商】回填手续费
// taobao.alitrip.seller.refund.fillfee
//
// 回填手续费
func TaobaoAlitripSellerRefundFillfee(clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerRefundFillfeeAPIRequest, resp *jipiao.TaobaoAlitripSellerRefundFillfeeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

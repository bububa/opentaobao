package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniDealerOdersRefundAddress 经销商查询逆向退货地址
// taobao.omni.dealer.oders.refund.address
//
// 经销商查询逆向退货地址
func TaobaoOmniDealerOdersRefundAddress(clt *core.SDKClient, req *omniorder.TaobaoOmniDealerOdersRefundAddressAPIRequest, resp *omniorder.TaobaoOmniDealerOdersRefundAddressAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

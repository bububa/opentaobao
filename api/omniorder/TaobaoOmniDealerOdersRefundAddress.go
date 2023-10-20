package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniDealerOdersRefundAddress 经销商查询逆向退货地址
// taobao.omni.dealer.oders.refund.address
//
// 经销商查询逆向退货地址
func TaobaoOmniDealerOdersRefundAddress(clt *core.SDKClient, req *omniorder.TaobaoOmniDealerOdersRefundAddressAPIRequest, session string) (*omniorder.TaobaoOmniDealerOdersRefundAddressAPIResponse, error) {
	var resp omniorder.TaobaoOmniDealerOdersRefundAddressAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

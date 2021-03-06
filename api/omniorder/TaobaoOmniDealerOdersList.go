package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniDealerOdersList 全渠道经销商订单列表
// taobao.omni.dealer.oders.list
//
// 全渠道经销商订单列表查询
func TaobaoOmniDealerOdersList(clt *core.SDKClient, req *omniorder.TaobaoOmniDealerOdersListAPIRequest, session string) (*omniorder.TaobaoOmniDealerOdersListAPIResponse, error) {
	var resp omniorder.TaobaoOmniDealerOdersListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

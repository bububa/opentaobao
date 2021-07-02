package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// AlibabaNlifeB2bTradeRefundList 获取采购退货单列表
// alibaba.nlife.b2b.trade.refund.list
//
// 获取采购退货单列表
func AlibabaNlifeB2bTradeRefundList(clt *core.SDKClient, req *nlife.AlibabaNlifeB2bTradeRefundListAPIRequest, session string) (*nlife.AlibabaNlifeB2bTradeRefundListAPIResponse, error) {
	var resp nlife.AlibabaNlifeB2bTradeRefundListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

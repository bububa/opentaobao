package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderDtdConsign 门店自送发货
// taobao.omniorder.dtd.consign
//
// 该接口触发门店自送发货，推进淘系订单状态为发货，为消费者发送核销码短信，并将物流信息写入订单
func TaobaoOmniorderDtdConsign(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniorderDtdConsignAPIRequest, resp *omniorder.TaobaoOmniorderDtdConsignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

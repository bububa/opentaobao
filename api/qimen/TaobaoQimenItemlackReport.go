package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenItemlackReport 发货单缺货通知接口
// taobao.qimen.itemlack.report
//
// WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP
func TaobaoQimenItemlackReport(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenItemlackReportAPIRequest, resp *qimen.TaobaoQimenItemlackReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

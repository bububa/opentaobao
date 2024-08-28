package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoRdcAligeniusWarehouseResendLogisticsMsgPost 补发单物流信息回传
// taobao.rdc.aligenius.warehouse.resend.logistics.msg.post
//
// 补发单erp物流信息回传平台
func TaobaoRdcAligeniusWarehouseResendLogisticsMsgPost(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest, resp *logistic.TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

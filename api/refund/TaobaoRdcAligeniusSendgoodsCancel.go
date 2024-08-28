package refund

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

// TaobaoRdcAligeniusSendgoodsCancel 取消发货
// taobao.rdc.aligenius.sendgoods.cancel
//
// 提供商家在仅退款中发送取消发货状态
func TaobaoRdcAligeniusSendgoodsCancel(ctx context.Context, clt *core.SDKClient, req *refund.TaobaoRdcAligeniusSendgoodsCancelAPIRequest, resp *refund.TaobaoRdcAligeniusSendgoodsCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoRdcAligeniusWarehouseResendUpdate 补发单状态回传
// taobao.rdc.aligenius.warehouse.resend.update
//
// 补发单状态回传接口
func TaobaoRdcAligeniusWarehouseResendUpdate(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest, resp *logistic.TaobaoRdcAligeniusWarehouseResendUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

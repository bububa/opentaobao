package lstlogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics"
)

// AlibabaLstShiporderCancel 零售通发货单取消
// alibaba.lst.shiporder.cancel
//
// 通过该接口可以取消零售通运保保发货单，并处理相关业务流程。
func AlibabaLstShiporderCancel(ctx context.Context, clt *core.SDKClient, req *lstlogistics.AlibabaLstShiporderCancelAPIRequest, resp *lstlogistics.AlibabaLstShiporderCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package lstlogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics"
)

// AlibabaLstShiporderCreate 零售通发货单创建
// alibaba.lst.shiporder.create
//
// 通过该接口可以创建零售通运保保发货单，并处理相关业务流程。
func AlibabaLstShiporderCreate(ctx context.Context, clt *core.SDKClient, req *lstlogistics.AlibabaLstShiporderCreateAPIRequest, resp *lstlogistics.AlibabaLstShiporderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

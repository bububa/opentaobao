package lstlogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics"
)

// AlibabaLstShiporderCancel 零售通发货单取消
// alibaba.lst.shiporder.cancel
//
// 通过该接口可以取消零售通运保保发货单，并处理相关业务流程。
func AlibabaLstShiporderCancel(clt *core.SDKClient, req *lstlogistics.AlibabaLstShiporderCancelAPIRequest, resp *lstlogistics.AlibabaLstShiporderCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

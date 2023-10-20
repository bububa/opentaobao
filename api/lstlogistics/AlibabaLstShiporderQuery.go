package lstlogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics"
)

// AlibabaLstShiporderQuery 零售通发货单查询
// alibaba.lst.shiporder.query
//
// 通过该接口可以查询零售通运保保发货单，并处理相关业务流程。
func AlibabaLstShiporderQuery(clt *core.SDKClient, req *lstlogistics.AlibabaLstShiporderQueryAPIRequest, resp *lstlogistics.AlibabaLstShiporderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

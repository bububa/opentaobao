package idleitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleitem"
)

// AlibabaIdleRecycleOrderGet 闲鱼回收订单查询V2
// alibaba.idle.recycle.order.get
//
// 闲鱼回收业务中,外部回收商作为交易上买家,闲鱼用户下单后,需要回收商主动拉取交易订单
func AlibabaIdleRecycleOrderGet(clt *core.SDKClient, req *idleitem.AlibabaIdleRecycleOrderGetAPIRequest, resp *idleitem.AlibabaIdleRecycleOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

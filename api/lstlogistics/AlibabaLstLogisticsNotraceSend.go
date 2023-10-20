package lstlogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics"
)

// AlibabaLstLogisticsNotraceSend 供应商-异云-无需物流发货
// alibaba.lst.logistics.notrace.send
//
// 异地云仓的订单，使用无需物流的发货方式来变更订单发货状态
func AlibabaLstLogisticsNotraceSend(clt *core.SDKClient, req *lstlogistics.AlibabaLstLogisticsNotraceSendAPIRequest, resp *lstlogistics.AlibabaLstLogisticsNotraceSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

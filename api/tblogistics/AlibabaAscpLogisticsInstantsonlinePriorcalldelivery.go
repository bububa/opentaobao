package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsInstantsonlinePriorcalldelivery 同城配送在线下单预询价
// alibaba.ascp.logistics.instantsonline.priorcalldelivery
//
// 同城配送在线下单预询价
func AlibabaAscpLogisticsInstantsonlinePriorcalldelivery(clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest, resp *tblogistics.AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

package lstlogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics"
)

// AlibabaLstLogisticsThirdpartSend 供应商-异云-使用第三方物流发货
// alibaba.lst.logistics.thirdpart.send
//
// 异地云仓的订单，使用第三方物流的发货方式来变更订单发货状态
func AlibabaLstLogisticsThirdpartSend(clt *core.SDKClient, req *lstlogistics.AlibabaLstLogisticsThirdpartSendAPIRequest, resp *lstlogistics.AlibabaLstLogisticsThirdpartSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

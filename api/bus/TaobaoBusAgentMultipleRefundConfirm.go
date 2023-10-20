package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusAgentMultipleRefundConfirm 综合交通多次退款接口
// taobao.bus.agent.multiple.refund.confirm
//
// 此接口支持多次按照单客进行多次退款操作，只进行退款操作。
func TaobaoBusAgentMultipleRefundConfirm(clt *core.SDKClient, req *bus.TaobaoBusAgentMultipleRefundConfirmAPIRequest, resp *bus.TaobaoBusAgentMultipleRefundConfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

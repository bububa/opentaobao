package cmns

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

// YunosServiceCmnsCoaMessageAck 消息回执查询
// yunos.service.cmns.coa.message.ack
//
// 第三方应用开发者调用此接口查询设备是否收到消息，只能查询此appKey床发的消息
func YunosServiceCmnsCoaMessageAck(clt *core.SDKClient, req *cmns.YunosServiceCmnsCoaMessageAckAPIRequest, resp *cmns.YunosServiceCmnsCoaMessageAckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentOrderHkAPIRequest
【国际机票】手工预定回填PNR API请求
taobao.alitrip.ie.agent.order.hk

代理商通过手工预定PNR，并回填。 */
type TaobaoAlitripIeAgentOrderHkAPIRequest struct {
	model.Params
	// 代理商ID
	_agentId int64
	// 回填pnr信息
	_writeBackPnrVO *IeWriteBackPnrVO
}

// New

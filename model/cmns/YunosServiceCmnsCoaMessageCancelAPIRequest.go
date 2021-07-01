package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosServiceCmnsCoaMessageCancelAPIRequest
CMNS消息撤回 API请求
yunos.service.cmns.coa.message.cancel

此接口用户撤回之前已经发出去的消息，根据消息ID撤回，只能撤回此appKey创建的消息。 */
type YunosServiceCmnsCoaMessageCancelAPIRequest struct {
	model.Params
	// 消息ID
	_mid int64
}

// New

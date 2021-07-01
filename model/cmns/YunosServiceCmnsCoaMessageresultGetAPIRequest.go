package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosServiceCmnsCoaMessageresultGetAPIRequest
CMNS消息发送到达查询 API请求
yunos.service.cmns.coa.messageresult.get

CMNS消息发送到达查询,根据消息ID查询，仅能查询该appKey所发送的消息 */
type YunosServiceCmnsCoaMessageresultGetAPIRequest struct {
	model.Params
	// 消息ID
	_mid int64
}

// New

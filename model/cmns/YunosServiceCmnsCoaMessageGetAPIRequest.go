package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosServiceCmnsCoaMessageGetAPIRequest
消息详情查询 API请求
yunos.service.cmns.coa.message.get

第三方应用开发者调用此接口查询消息详情，只能查询此appKey发的消息 */
type YunosServiceCmnsCoaMessageGetAPIRequest struct {
	model.Params
	// 消息id
	_mid int64
}

// New

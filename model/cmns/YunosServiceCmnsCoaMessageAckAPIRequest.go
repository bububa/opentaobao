package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosServiceCmnsCoaMessageAckAPIRequest
消息回执查询 API请求
yunos.service.cmns.coa.message.ack

第三方应用开发者调用此接口查询设备是否收到消息，只能查询此appKey床发的消息 */
type YunosServiceCmnsCoaMessageAckAPIRequest struct {
	model.Params
	// 设备唯一值deviceToken
	_deviceToken string
	// 设备imei
	_imei string
	// 设备uuid
	_uuid string
	// 消息ID
	_mid int64
}

// New

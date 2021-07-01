package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosServiceCmnsCoaMessagePushAPIRequest
消息推送接口 API请求
yunos.service.cmns.coa.message.push

调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。 */
type YunosServiceCmnsCoaMessagePushAPIRequest struct {
	model.Params
	// 消息推送请求对象
	_pushRequest *PushRequest
}

// New

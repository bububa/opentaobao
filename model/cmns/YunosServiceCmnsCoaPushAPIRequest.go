package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosServiceCmnsCoaPushAPIRequest
消息推送接口 API请求
yunos.service.cmns.coa.push

调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。 */
type YunosServiceCmnsCoaPushAPIRequest struct {
	model.Params
	// 消息结构对象
	_msgObj *CmnsMessage
}

// New

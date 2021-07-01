package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractActivityUnregisterAPIRequest
ISV互动应用活动关闭服务 API请求
alibaba.interact.activity.unregister

卖家在ISV互动应用中设置的活动主动关闭的服务 */
type AlibabaInteractActivityUnregisterAPIRequest struct {
	model.Params
	// 互动活动ID
	_bizId string
}

// New

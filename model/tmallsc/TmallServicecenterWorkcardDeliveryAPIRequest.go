package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardDeliveryAPIRequest
开始配送工单 API请求
tmall.servicecenter.workcard.delivery

服务商调用该接口通知天猫服务平台服务商工人已开始配送工单 */
type TmallServicecenterWorkcardDeliveryAPIRequest struct {
	model.Params
	// 工单配送请求参数
	_identifyTaskDeliveryRequest *IdentifyTaskDeliveryRequest
}

// New

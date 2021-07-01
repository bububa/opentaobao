package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardVirtualphoneBindAPIRequest
工单维度虚拟中间号绑定 API请求
tmall.servicecenter.workcard.virtualphone.bind

服务供应链洗护服务ERP项目中，客服呼叫消费者的功能。
叫消费者的手机号虚拟号码给到客服。 */
type TmallServicecenterWorkcardVirtualphoneBindAPIRequest struct {
	model.Params
	// 绑定阿里通讯号入参
	_workcardRequest *WorkcardBaseRequest
}

// New

package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallAliautoAutofinanceCreditReceiveAPIRequest
接收授信结果通知 API请求
tmall.aliauto.autofinance.credit.receive

天猫汽车的金融业务场景中，需要接收外部ISV对用户授信申请的通知结果. */
type TmallAliautoAutofinanceCreditReceiveAPIRequest struct {
	model.Params
	// 授信通知描述
	_creditReceiveDto *CreditReceiveDto
}

// New

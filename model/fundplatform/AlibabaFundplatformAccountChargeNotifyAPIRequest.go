package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFundplatformAccountChargeNotifyAPIRequest
账户充值成功通知 API请求
alibaba.fundplatform.account.charge.notify

通知外部业务方充值成功 */
type AlibabaFundplatformAccountChargeNotifyAPIRequest struct {
	model.Params
	// 入参对象
	_request *AlibabaFundplatformAccountChargeNotifyStruct
}

// New

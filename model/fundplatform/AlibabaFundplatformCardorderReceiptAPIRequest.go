package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFundplatformCardorderReceiptAPIRequest
通知确认收货 API请求
alibaba.fundplatform.cardorder.receipt

告知卡商这一批储值卡已经被用户确认收货 */
type AlibabaFundplatformCardorderReceiptAPIRequest struct {
	model.Params
	// 通知制卡成功的子卡子单号
	_cardOrderId int64
	// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
	_ownSign string
}

// New

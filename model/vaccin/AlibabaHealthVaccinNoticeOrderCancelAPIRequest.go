package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinNoticeOrderCancelAPIRequest
福州疫苗取消预约 API请求
alibaba.health.vaccin.notice.order.cancel

福州疫苗用户取消预约接口 */
type AlibabaHealthVaccinNoticeOrderCancelAPIRequest struct {
	model.Params
	// 支付宝用户id
	_alipayUserId string
	// 预约id
	_orderId string
}

// New

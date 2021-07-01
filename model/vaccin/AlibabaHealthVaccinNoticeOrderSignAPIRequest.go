package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinNoticeOrderSignAPIRequest
福州疫苗签到成功通知 API请求
alibaba.health.vaccin.notice.order.sign

福州疫苗用户签到成功记录 */
type AlibabaHealthVaccinNoticeOrderSignAPIRequest struct {
	model.Params
	// 支付宝用户id
	_alipayUserId string
	// 预约id
	_orderId string
}

// New

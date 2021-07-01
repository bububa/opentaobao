package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinNoticeReplantRemindAPIRequest
支付宝疫苗补种提醒信息 API请求
alibaba.health.vaccin.notice.replant.remind

支付宝疫苗补种提醒 */
type AlibabaHealthVaccinNoticeReplantRemindAPIRequest struct {
	model.Params
	// 支付宝ID
	_alipayUserId string
	// 针次
	_theTimes string
	// 预约id
	_orderId string
}

// New

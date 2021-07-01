package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinNoticeOrderCompleteAPIRequest
疫苗接种完成(带支付宝提醒) API请求
alibaba.health.vaccin.notice.order.complete

用户到店完成接种,ISV感知通知阿里健康完成接种,并通知用户! */
type AlibabaHealthVaccinNoticeOrderCompleteAPIRequest struct {
	model.Params
	// 支付宝唯一标识
	_alipayUserId string
	// 在ISV预约单唯一标识
	_orderId string
}

// New

package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinNoticeMissRemindAPIRequest
疫苗漏种提醒 API请求
alibaba.health.vaccin.notice.miss.remind

医生消息提醒适龄儿童按计划接种 */
type AlibabaHealthVaccinNoticeMissRemindAPIRequest struct {
	model.Params
	// 432421
	_alipayUserId string
	// 多个疫苗英文逗号分隔
	_vaccineName string
	// 2019-02-08 严格按照
	_reserveDate string
	// 姓名
	_name string
	// 点击提醒消息的跳转链接
	_url string
}

// New

package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest
疫苗预约时间段提醒 API请求
alibaba.health.vaccin.notice.timebucket.remind

疫苗预约时间段提醒 */
type AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest struct {
	model.Params
	// 432421
	_alipayUserId string
	// 疫苗名称
	_vaccineName string
	// 预约日期：2019-02-08 严格按照
	_reserveDate string
	// 接种人姓名
	_name string
	// 针次
	_theTimes string
	// 接种点名称（通知方）
	_povStoreName string
	// 可预约时段
	_reserveTime string
	// 用户授权的手机号
	_mobile string
}

// New

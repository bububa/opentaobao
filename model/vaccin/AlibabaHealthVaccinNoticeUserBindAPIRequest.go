package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinNoticeUserBindAPIRequest
支付宝疫苗绑定接种人 API请求
alibaba.health.vaccin.notice.user.bind

支付宝疫苗绑定接种人 */
type AlibabaHealthVaccinNoticeUserBindAPIRequest struct {
	model.Params
	// 支付宝ID
	_alipayUserId string
	// 绑定人信息list
	_bindUsers []AlipayVaccineUserBindDto
	// ISV 侧用户 ID
	_outerUserId string
	// 联系电话
	_mobile string
	// 用户来源：alipay或yl
	_appChannel string
}

// New

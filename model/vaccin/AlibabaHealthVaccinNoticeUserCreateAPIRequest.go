package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinNoticeUserCreateAPIRequest
支付宝医疗健康疫苗用户创建 API请求
alibaba.health.vaccin.notice.user.create

支付宝医疗健康疫苗用户创建 */
type AlibabaHealthVaccinNoticeUserCreateAPIRequest struct {
	model.Params
	// 支付宝用户id
	_aliPayUserId string
	// 外部渠道用户id
	_outerUserId string
	// 用户电话号码
	_mobile string
}

// New

package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinVaccinateCompleteAPIRequest
接种完成反馈接口 API请求
alibaba.health.vaccin.vaccinate.complete

ISV 将用户完成接种的疫苗同步给免疫规划中心 */
type AlibabaHealthVaccinVaccinateCompleteAPIRequest struct {
	model.Params
	// 支付宝用户 ID
	_alipayUserId string
	// ISV 侧用户 ID
	_isvUserId string
	// 订单 ID
	_orderId string
	// 接种人姓名
	_name string
	// 联系电话
	_mobile string
	// 接种日期
	_vaccinateDate string
	// 接种时间
	_vaccinateTime string
	// 接种的疫苗信息
	_vaccineList []VaccineInfo
}

// New

package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinNoticeSendAPIRequest
发送消息提醒 API请求
alibaba.health.vaccin.notice.send

ISV 通过免疫规划中心给用户发送短信或者支付宝 PUSH 提醒。 */
type AlibabaHealthVaccinNoticeSendAPIRequest struct {
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
	// 接种的疫苗信息
	_vaccineList []VaccineInfo
	// 接种人性别:1=男,2=女
	_sex int64
	// 接种人出生日期
	_birthday string
	// 接种点编码
	_povNo string
	// 接种点名称
	_povName string
	// 接种点地址
	_address string
	// 省份名称
	_province string
	// 城市名称
	_city string
	// 区名称
	_area string
	// 预约日期
	_reserveDate string
	// 预约时间段
	_reserveTime string
	// 消息提醒类型：1=疫苗下一针预约提醒  2=预约成功提醒  3=接种提醒  4=补种提醒
	_messageType int64
	// 用户入口，支付宝或医鹿，alipay或yl
	_appChannel string
}

// New

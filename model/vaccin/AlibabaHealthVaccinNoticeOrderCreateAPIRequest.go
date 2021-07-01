package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinNoticeOrderCreateAPIRequest
支付宝医疗健康疫苗预约创建 API请求
alibaba.health.vaccin.notice.order.create

支付宝医疗健康疫苗预约创建 */
type AlibabaHealthVaccinNoticeOrderCreateAPIRequest struct {
	model.Params
	// 预约人性别(1男2女)
	_sex int64
	// 年龄
	_age int64
	// 预约日期
	_reserveDate string
	// 支付宝用户id
	_alipayUserId string
	// 外部渠道用户id
	_outerUserId string
	// 预约id
	_orderId string
	// 手机号码
	_mobile string
	// 接种人姓名
	_name string
	// 接种点地址
	_address string
	// 接种点名称
	_povStoreName string
	// 预约时间
	_reserveTime string
	// 疫苗信息
	_vaccineInfo string
	// 年龄类型(1-宝宝2-成人)
	_ageType int64
	// 支付宝消息通知跳转订单详情链接
	_orderDetailUrl string
	// 地区名字
	_area string
	// 城市名字
	_city string
	// 省份名字
	_province string
}

// New

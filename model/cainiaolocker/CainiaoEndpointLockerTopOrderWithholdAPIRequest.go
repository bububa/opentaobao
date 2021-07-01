package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoEndpointLockerTopOrderWithholdAPIRequest
代扣支付 API请求
cainiao.endpoint.locker.top.order.withhold

提供代扣，允许有一笔欠款。 */
type CainiaoEndpointLockerTopOrderWithholdAPIRequest struct {
	model.Params
	// 柜子公司编码
	_companyCode string
	// 柜子id
	_guiId string
	// 订单类型(0-取件业务，1-寄件业务，2-派样业务)
	_orderType int64
	// 开放用户id
	_openUserId string
	// 代扣金额（全额），单位：分
	_totalFee int64
	// 扩展字段
	_extra string
	// 柜子订单编码
	_orderCode string
	// 运单号
	_mailNo string
}

// New

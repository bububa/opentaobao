package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoNbaddAppointdeliverGetconsigninfoAPIRequest
获取支持定时派送服务发货信息 API请求
cainiao.nbadd.appointdeliver.getconsigninfo

获取支持定时派送服务发货信息 */
type CainiaoNbaddAppointdeliverGetconsigninfoAPIRequest struct {
	model.Params
	// 淘宝交易订单id
	_tradeOrderId int64
}

// New

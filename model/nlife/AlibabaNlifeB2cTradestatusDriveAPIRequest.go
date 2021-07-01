package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeB2cTradestatusDriveAPIRequest
b2c订单状态驱动 API请求
alibaba.nlife.b2c.tradestatus.drive

用于驱动零售+订单状态 */
type AlibabaNlifeB2cTradestatusDriveAPIRequest struct {
	model.Params
	// 零售门店在零售+平台的ID
	_storeId string
	// APP:是指线上销售应用，POS:是指现场收银应用
	_channel string
	// 对零售+为外部订单号，对业务方为订单号
	_outTradeNo string
	// 零售+平台订单号，和out_trade_no不能同时为空
	_tradeNo string
	// 接口类型：CONFIRM（收货）DELIVER（发货）
	_action string
	// 货流信息
	_logisticsInfo *LogisticsInfo
	// 扩展参数 JSON格式
	_extendParams string
}

// New

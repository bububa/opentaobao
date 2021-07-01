package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallDeviceTradePrecreateAPIRequest
智能硬件上预创建天猫订单 API请求
tmall.device.trade.precreate

智能硬件上预创建天猫订单。
1，use_open_price不再需要传入，使用unit_price传入单价。
2，订单默认5分钟自动关闭，没有付款的订单在手机淘宝不可见。
3，同一个码只运行一个用户扫码，多个用户扫一个码会报错 订单不存在。 */
type TmallDeviceTradePrecreateAPIRequest struct {
	model.Params
	// 交易类型。1，售卖。2，派样
	_type int64
	// 商品列表
	_itemList []TradeItemDo
	// 设备业务编码
	_deviceCode string
	// 外部订单ID，用户下保证唯一。最大长度22
	_outTradeId string
	// 回调地址，当订单创建，付款成功后，会收到回调。必须是https地址，HTTP响应必须包含success，否则系统会重发
	_callbackUrl string
	// 导购员手机号码
	_baPhone string
	// 导购员所属门店
	_baStoreId int64
}

// New

package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiOrderNotifyAPIRequest
状态通知 API请求
alibaba.happytrip.taxi.order.notify

当订单发生变化是供应商通过状态通知API通知欢行，欢行获取最新的订单详情和状态进行业务处理。 */
type AlibabaHappytripTaxiOrderNotifyAPIRequest struct {
	model.Params
	// 返回自 1970 年 1 月 1 日 00:00:00 GMT 以来此 Date 对象表示的毫秒数
	_time int64
	// 通知类型: 1-订单中间状态流转 2-订单终态通知 3-支付确认通知 4-订单退款通知 5-订单改价通知 6-客服关单通知。参考：https://open-hatrip.alibaba.com/doc/car/order_status_callback.html
	_notifyType int64
	// 通知说明
	_notifyDesc string
	// 订单id
	_orderId string
}

// New

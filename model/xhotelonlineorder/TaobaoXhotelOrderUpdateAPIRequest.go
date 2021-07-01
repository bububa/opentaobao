package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderUpdateAPIRequest
酒店订单发货接口 API请求
taobao.xhotel.order.update

卖家确认订单或者取消订单，适用于预付、面付、信用住订单 */
type TaobaoXhotelOrderUpdateAPIRequest struct {
	model.Params
	// 订单号
	_tid int64
	// 操作的类型：1.确认无房（取消预订，710发送短信提醒买家申请退款）2.确认预订 3.入住 4.离店 5.noshow 6.关单
	_optType int64
	// 是否把代理直签的订单同步到酒店，Y为同步，N不同步
	_syncToHotel string
	// 退款费用
	_refundFee int64
	// 取消类型，6 代表的是用户取消，reasonType=7代表的是小二协商
	_reasonType int64
	// 开票金额
	_invoiceAmount int64
}

// New

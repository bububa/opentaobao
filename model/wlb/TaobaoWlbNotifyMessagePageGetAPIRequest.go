package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbNotifyMessagePageGetAPIRequest
物流宝通知消息查询接口 API请求
taobao.wlb.notify.message.page.get

物流宝提供的消息通知查询接口，消息内容包括;出入库单不一致消息，取消订单成功消息，盘点单消息 */
type TaobaoWlbNotifyMessagePageGetAPIRequest struct {
	model.Params
	// 通知消息编码： STOCK_IN_NOT_CONSISTENT---入库单不一致 CANCEL_ORDER_SUCCESS---取消订单成功 INVENTORY_CHECK---盘点 CANCEL_ORDER_FAILURE---取消订单失败 ORDER_REJECT--wms拒单 ORDER_CONFIRMED--订单处理成功
	_msgCode string
	// 分页查询页数
	_pageNo int64
	// 分页查询的每页页数
	_pageSize int64
	// 记录开始时间
	_startDate string
	// 记录截至时间
	_endDate string
	// 消息状态： 不需要确认：NO_NEED_CONFIRM 已确认：CONFIRMED 待确认：TO_BE_CONFIRM
	_status string
}

// New

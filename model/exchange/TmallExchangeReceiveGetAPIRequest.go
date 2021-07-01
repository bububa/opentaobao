package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallExchangeReceiveGetAPIRequest
卖家查询换货列表 API请求
tmall.exchange.receive.get

卖家查询换货列表 */
type TmallExchangeReceiveGetAPIRequest struct {
	model.Params
	// 查询修改时间段的结束时间点
	_endGmtModifedTime string
	// 查询修改时间段的开始时间点
	_startGmtModifiedTime string
	// 快递单号
	_logisticNo string
	// 买家昵称
	_buyerNick string
	// 查询申请时间段的开始时间点
	_startCreatedTime string
	// 返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick
	_fields []string
	// 每页条数
	_pageSize int64
	// 换货状态，具体包括：换货待处理(1), 待买家退货(2), 买家已退货，待收货(3),  换货关闭(4), 换货成功(5), 待买家修改(6), 待发出换货商品(12), 待买家收货(13), 请退款(14)
	_disputeStatusArray []int64
	// 查询申请时间段的结束时间点
	_endCreatedTime string
	// 买家id
	_buyerId int64
	// 退款单号ID列表，最多只能输入20个id
	_refundIdArray []int64
	// 页码
	_pageNo int64
	// 正向订单号
	_bizOrderId int64
}

// New

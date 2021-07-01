package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripSellerRefundorderlistFetchAPIRequest
【机票代理商】——退票订单列表提取 API请求
taobao.alitrip.seller.refundorderlist.fetch

代理商纬度退票订单列表提取 */
type TaobaoAlitripSellerRefundorderlistFetchAPIRequest struct {
	model.Params
	// 提取数据的开始时间
	_startDate string
	// 1：初始，2：接受，3：确认，4：失败，5：买家处理，6：卖家处理，7：等待小二回填，8：退款成功
	_status int64
	// 提取数据的结束时间
	_endDate string
}

// New

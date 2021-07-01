package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripSellerRefundSearchAPIRequest
【机票代理商】退票申请单列表 API请求
taobao.alitrip.seller.refund.search

查询退票申请单列表 */
type TaobaoAlitripSellerRefundSearchAPIRequest struct {
	model.Params
	// 结束时间
	_endTime string
	// 开始时间
	_startTime string
	// 申请单状态（如果为空查询所有状态，1初始 2接受 3确认 4失败 5买家处理 6卖家处理 7等待小二回填 8退款成功）
	_status int64
}

// New

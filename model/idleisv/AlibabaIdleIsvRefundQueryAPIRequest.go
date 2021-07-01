package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleIsvRefundQueryAPIRequest
闲鱼已验货交易订单退款信息查询 API请求
alibaba.idle.isv.refund.query

闲鱼服务商交易订单退款信息查询-单个退款查询 */
type AlibabaIdleIsvRefundQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramAppraiseIsvOrderQuery *AppraiseIsvOrderQuery
}

// New

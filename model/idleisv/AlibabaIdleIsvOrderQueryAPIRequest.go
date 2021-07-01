package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleIsvOrderQueryAPIRequest
闲鱼已验货订单查询 API请求
alibaba.idle.isv.order.query

服务商ISV，根据订单号，查询闲鱼订单信息 */
type AlibabaIdleIsvOrderQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramAppraiseIsvOrderQuery *AppraiseIsvOrderQuery
}

// New

package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosOrderqsMisbigposOrderQueryAPIRequest
大pos新选单退 API请求
alibaba.mos.orderqs.misbigpos.order.query

大pos新选单退 */
type AlibabaMosOrderqsMisbigposOrderQueryAPIRequest struct {
	model.Params
	// 外部门店号
	_storeNo string
	// 基本信息获取参数
	_queryBaseData bool
	// 小票号
	_receiptNo string
	// 券扩展数据获取
	_queryCouponExtern bool
}

// New

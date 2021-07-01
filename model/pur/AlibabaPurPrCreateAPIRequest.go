package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPurPrCreateAPIRequest
下pr单 API请求
alibaba.pur.pr.create

下pr单 */
type AlibabaPurPrCreateAPIRequest struct {
	model.Params
	// 订单信息
	_purReq *MallReceivePrRequest
}

// New

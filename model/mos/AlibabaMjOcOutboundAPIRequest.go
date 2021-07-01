package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjOcOutboundAPIRequest
零售商品发货 API请求
alibaba.mj.oc.outbound

用于接收发货的数据 */
type AlibabaMjOcOutboundAPIRequest struct {
	model.Params
	// 发货信息
	_goodsOutbound *GoodsOutboundDto
}

// New

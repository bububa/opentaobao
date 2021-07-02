package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcOutboundAPIRequest 零售商品发货 API请求
// alibaba.mj.oc.outbound
//
// 用于接收发货的数据
type AlibabaMjOcOutboundAPIRequest struct {
	model.Params
	// 发货信息
	_goodsOutbound *GoodsOutboundDto
}

// NewAlibabaMjOcOutboundRequest 初始化AlibabaMjOcOutboundAPIRequest对象
func NewAlibabaMjOcOutboundRequest() *AlibabaMjOcOutboundAPIRequest {
	return &AlibabaMjOcOutboundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjOcOutboundAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.outbound"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjOcOutboundAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is GoodsOutbound Setter
// 发货信息
func (r *AlibabaMjOcOutboundAPIRequest) SetGoodsOutbound(_goodsOutbound *GoodsOutboundDto) error {
	r._goodsOutbound = _goodsOutbound
	r.Set("goods_outbound", _goodsOutbound)
	return nil
}

// Get GoodsOutbound Getter
func (r AlibabaMjOcOutboundAPIRequest) GetGoodsOutbound() *GoodsOutboundDto {
	return r._goodsOutbound
}

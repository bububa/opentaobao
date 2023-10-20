package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjocoutboundAPIRequest 零售商品发货 API请求
// alibaba.mj.oc.outbound
//
// 用于接收发货的数据
type AlibabamjocoutboundAPIRequest struct {
	model.Params
	// 发货信息
	_goodsOutbound *GoodsOutboundDto
}

// NewAlibabamjocoutboundRequest 初始化AlibabamjocoutboundAPIRequest对象
func NewAlibabamjocoutboundRequest() *AlibabamjocoutboundAPIRequest {
	return &AlibabamjocoutboundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamjocoutboundAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.outbound"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamjocoutboundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamjocoutboundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsOutbound is GoodsOutbound Setter
// 发货信息
func (r *AlibabamjocoutboundAPIRequest) SetGoodsOutbound(_goodsOutbound *GoodsOutboundDto) error {
	r._goodsOutbound = _goodsOutbound
	r.Set("goods_outbound", _goodsOutbound)
	return nil
}

// GetGoodsOutbound GoodsOutbound Getter
func (r AlibabamjocoutboundAPIRequest) GetGoodsOutbound() *GoodsOutboundDto {
	return r._goodsOutbound
}

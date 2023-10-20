package mos

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMjOcOutboundAPIRequest) Reset() {
	r._goodsOutbound = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjOcOutboundAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.outbound"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjOcOutboundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMjOcOutboundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsOutbound is GoodsOutbound Setter
// 发货信息
func (r *AlibabaMjOcOutboundAPIRequest) SetGoodsOutbound(_goodsOutbound *GoodsOutboundDto) error {
	r._goodsOutbound = _goodsOutbound
	r.Set("goods_outbound", _goodsOutbound)
	return nil
}

// GetGoodsOutbound GoodsOutbound Getter
func (r AlibabaMjOcOutboundAPIRequest) GetGoodsOutbound() *GoodsOutboundDto {
	return r._goodsOutbound
}

var poolAlibabaMjOcOutboundAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMjOcOutboundRequest()
	},
}

// GetAlibabaMjOcOutboundRequest 从 sync.Pool 获取 AlibabaMjOcOutboundAPIRequest
func GetAlibabaMjOcOutboundAPIRequest() *AlibabaMjOcOutboundAPIRequest {
	return poolAlibabaMjOcOutboundAPIRequest.Get().(*AlibabaMjOcOutboundAPIRequest)
}

// ReleaseAlibabaMjOcOutboundAPIRequest 将 AlibabaMjOcOutboundAPIRequest 放入 sync.Pool
func ReleaseAlibabaMjOcOutboundAPIRequest(v *AlibabaMjOcOutboundAPIRequest) {
	v.Reset()
	poolAlibabaMjOcOutboundAPIRequest.Put(v)
}

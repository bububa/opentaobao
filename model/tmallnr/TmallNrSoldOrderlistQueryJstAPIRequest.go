package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrsoldorderlistqueryjstAPIRequest 已入塔商家查询订单列表 API请求
// tmall.nr.sold.orderlist.query.jst
//
// 该服务用于已入聚石塔的商家，获取订单列表信息；
type TmallnrsoldorderlistqueryjstAPIRequest struct {
	model.Params
	// 入参对象
	_param0 *NrTimingOrderSoldQueryReqDto
}

// NewTmallnrsoldorderlistqueryjstRequest 初始化TmallnrsoldorderlistqueryjstAPIRequest对象
func NewTmallnrsoldorderlistqueryjstRequest() *TmallnrsoldorderlistqueryjstAPIRequest {
	return &TmallnrsoldorderlistqueryjstAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrsoldorderlistqueryjstAPIRequest) GetApiMethodName() string {
	return "tmall.nr.sold.orderlist.query.jst"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrsoldorderlistqueryjstAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrsoldorderlistqueryjstAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参对象
func (r *TmallnrsoldorderlistqueryjstAPIRequest) SetParam0(_param0 *NrTimingOrderSoldQueryReqDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallnrsoldorderlistqueryjstAPIRequest) GetParam0() *NrTimingOrderSoldQueryReqDto {
	return r._param0
}

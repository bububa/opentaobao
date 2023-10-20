package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchannelsupplierproductgoodsbindAPIRequest 渠道产品与货品绑定接口 API请求
// alibaba.ascp.channel.supplier.product.goods.bind
//
// 渠道产品与货品绑定接口
type AlibabaascpchannelsupplierproductgoodsbindAPIRequest struct {
	model.Params
	// 请求
	_topBindProductGoodsRequest *TopBindProductGoodsRequest
}

// NewAlibabaascpchannelsupplierproductgoodsbindRequest 初始化AlibabaascpchannelsupplierproductgoodsbindAPIRequest对象
func NewAlibabaascpchannelsupplierproductgoodsbindRequest() *AlibabaascpchannelsupplierproductgoodsbindAPIRequest {
	return &AlibabaascpchannelsupplierproductgoodsbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchannelsupplierproductgoodsbindAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.supplier.product.goods.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchannelsupplierproductgoodsbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchannelsupplierproductgoodsbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopBindProductGoodsRequest is TopBindProductGoodsRequest Setter
// 请求
func (r *AlibabaascpchannelsupplierproductgoodsbindAPIRequest) SetTopBindProductGoodsRequest(_topBindProductGoodsRequest *TopBindProductGoodsRequest) error {
	r._topBindProductGoodsRequest = _topBindProductGoodsRequest
	r.Set("top_bind_product_goods_request", _topBindProductGoodsRequest)
	return nil
}

// GetTopBindProductGoodsRequest TopBindProductGoodsRequest Getter
func (r AlibabaascpchannelsupplierproductgoodsbindAPIRequest) GetTopBindProductGoodsRequest() *TopBindProductGoodsRequest {
	return r._topBindProductGoodsRequest
}

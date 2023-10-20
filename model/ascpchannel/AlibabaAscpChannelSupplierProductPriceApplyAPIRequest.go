package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchannelsupplierproductpriceapplyAPIRequest 供应商设置渠道产品价格 API请求
// alibaba.ascp.channel.supplier.product.price.apply
//
// 供应商设置渠道产品价格
type AlibabaascpchannelsupplierproductpriceapplyAPIRequest struct {
	model.Params
	// 请求参数
	_topPriceApplyCreateRequest *TopPriceApplyCreateRequest
}

// NewAlibabaascpchannelsupplierproductpriceapplyRequest 初始化AlibabaascpchannelsupplierproductpriceapplyAPIRequest对象
func NewAlibabaascpchannelsupplierproductpriceapplyRequest() *AlibabaascpchannelsupplierproductpriceapplyAPIRequest {
	return &AlibabaascpchannelsupplierproductpriceapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchannelsupplierproductpriceapplyAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.supplier.product.price.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchannelsupplierproductpriceapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchannelsupplierproductpriceapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopPriceApplyCreateRequest is TopPriceApplyCreateRequest Setter
// 请求参数
func (r *AlibabaascpchannelsupplierproductpriceapplyAPIRequest) SetTopPriceApplyCreateRequest(_topPriceApplyCreateRequest *TopPriceApplyCreateRequest) error {
	r._topPriceApplyCreateRequest = _topPriceApplyCreateRequest
	r.Set("top_price_apply_create_request", _topPriceApplyCreateRequest)
	return nil
}

// GetTopPriceApplyCreateRequest TopPriceApplyCreateRequest Getter
func (r AlibabaascpchannelsupplierproductpriceapplyAPIRequest) GetTopPriceApplyCreateRequest() *TopPriceApplyCreateRequest {
	return r._topPriceApplyCreateRequest
}

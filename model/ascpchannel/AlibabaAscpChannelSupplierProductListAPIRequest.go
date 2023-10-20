package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchannelsupplierproductlistAPIRequest 供应商渠道产品列表查询 API请求
// alibaba.ascp.channel.supplier.product.list
//
// 供应商查询渠道产品列表
type AlibabaascpchannelsupplierproductlistAPIRequest struct {
	model.Params
	// 请求参数
	_param *ProductListQueryRequestForSupplier
}

// NewAlibabaascpchannelsupplierproductlistRequest 初始化AlibabaascpchannelsupplierproductlistAPIRequest对象
func NewAlibabaascpchannelsupplierproductlistRequest() *AlibabaascpchannelsupplierproductlistAPIRequest {
	return &AlibabaascpchannelsupplierproductlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchannelsupplierproductlistAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.supplier.product.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchannelsupplierproductlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchannelsupplierproductlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求参数
func (r *AlibabaascpchannelsupplierproductlistAPIRequest) SetParam(_param *ProductListQueryRequestForSupplier) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaascpchannelsupplierproductlistAPIRequest) GetParam() *ProductListQueryRequestForSupplier {
	return r._param
}

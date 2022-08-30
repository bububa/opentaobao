package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelSupplierProductListAPIRequest 供应商渠道产品列表查询 API请求
// alibaba.ascp.channel.supplier.product.list
//
// 供应商查询渠道产品列表
type AlibabaAscpChannelSupplierProductListAPIRequest struct {
	model.Params
	// 请求参数
	_param *ProductListQueryRequestForSupplier
}

// NewAlibabaAscpChannelSupplierProductListRequest 初始化AlibabaAscpChannelSupplierProductListAPIRequest对象
func NewAlibabaAscpChannelSupplierProductListRequest() *AlibabaAscpChannelSupplierProductListAPIRequest {
	return &AlibabaAscpChannelSupplierProductListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelSupplierProductListAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.supplier.product.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelSupplierProductListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 请求参数
func (r *AlibabaAscpChannelSupplierProductListAPIRequest) SetParam(_param *ProductListQueryRequestForSupplier) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaAscpChannelSupplierProductListAPIRequest) GetParam() *ProductListQueryRequestForSupplier {
	return r._param
}

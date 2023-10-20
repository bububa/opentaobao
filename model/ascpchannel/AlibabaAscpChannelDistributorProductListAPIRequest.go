package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchanneldistributorproductlistAPIRequest 供应链渠道中心淘外分销品批量查询(分销商专用) API请求
// alibaba.ascp.channel.distributor.product.list
//
// 此api为淘外分销的品批量查询标准api，淘外分销商专用
type AlibabaascpchanneldistributorproductlistAPIRequest struct {
	model.Params
	// 列表请求
	_productListRequest *Productlistrequest
}

// NewAlibabaascpchanneldistributorproductlistRequest 初始化AlibabaascpchanneldistributorproductlistAPIRequest对象
func NewAlibabaascpchanneldistributorproductlistRequest() *AlibabaascpchanneldistributorproductlistAPIRequest {
	return &AlibabaascpchanneldistributorproductlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchanneldistributorproductlistAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.distributor.product.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchanneldistributorproductlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchanneldistributorproductlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductListRequest is ProductListRequest Setter
// 列表请求
func (r *AlibabaascpchanneldistributorproductlistAPIRequest) SetProductListRequest(_productListRequest *Productlistrequest) error {
	r._productListRequest = _productListRequest
	r.Set("product_list_request", _productListRequest)
	return nil
}

// GetProductListRequest ProductListRequest Getter
func (r AlibabaascpchanneldistributorproductlistAPIRequest) GetProductListRequest() *Productlistrequest {
	return r._productListRequest
}

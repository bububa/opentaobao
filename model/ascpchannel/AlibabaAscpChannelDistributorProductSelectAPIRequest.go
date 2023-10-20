package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchanneldistributorproductselectAPIRequest 供应链渠道中心品的选品接口（淘外分销商专用） API请求
// alibaba.ascp.channel.distributor.product.select
//
// 此api为淘外分销的品的选品标准api，淘外分销商专用
type AlibabaascpchanneldistributorproductselectAPIRequest struct {
	model.Params
	// 选品请求
	_selectProductRequest *ProductLinkRequest
}

// NewAlibabaascpchanneldistributorproductselectRequest 初始化AlibabaascpchanneldistributorproductselectAPIRequest对象
func NewAlibabaascpchanneldistributorproductselectRequest() *AlibabaascpchanneldistributorproductselectAPIRequest {
	return &AlibabaascpchanneldistributorproductselectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchanneldistributorproductselectAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.distributor.product.select"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchanneldistributorproductselectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchanneldistributorproductselectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSelectProductRequest is SelectProductRequest Setter
// 选品请求
func (r *AlibabaascpchanneldistributorproductselectAPIRequest) SetSelectProductRequest(_selectProductRequest *ProductLinkRequest) error {
	r._selectProductRequest = _selectProductRequest
	r.Set("select_product_request", _selectProductRequest)
	return nil
}

// GetSelectProductRequest SelectProductRequest Getter
func (r AlibabaascpchanneldistributorproductselectAPIRequest) GetSelectProductRequest() *ProductLinkRequest {
	return r._selectProductRequest
}

package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelDistributorProductSelectAPIRequest 供应链渠道中心品的选品接口（淘外分销商专用） API请求
// alibaba.ascp.channel.distributor.product.select
//
// 此api为淘外分销的品的选品标准api，淘外分销商专用
type AlibabaAscpChannelDistributorProductSelectAPIRequest struct {
	model.Params
	// 选品请求
	_selectProductRequest *ProductLinkRequest
}

// NewAlibabaAscpChannelDistributorProductSelectRequest 初始化AlibabaAscpChannelDistributorProductSelectAPIRequest对象
func NewAlibabaAscpChannelDistributorProductSelectRequest() *AlibabaAscpChannelDistributorProductSelectAPIRequest {
	return &AlibabaAscpChannelDistributorProductSelectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelDistributorProductSelectAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.distributor.product.select"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelDistributorProductSelectAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSelectProductRequest is SelectProductRequest Setter
// 选品请求
func (r *AlibabaAscpChannelDistributorProductSelectAPIRequest) SetSelectProductRequest(_selectProductRequest *ProductLinkRequest) error {
	r._selectProductRequest = _selectProductRequest
	r.Set("select_product_request", _selectProductRequest)
	return nil
}

// GetSelectProductRequest SelectProductRequest Getter
func (r AlibabaAscpChannelDistributorProductSelectAPIRequest) GetSelectProductRequest() *ProductLinkRequest {
	return r._selectProductRequest
}

package ascpchannel

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpChannelDistributorProductSelectAPIRequest) Reset() {
	r._selectProductRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelDistributorProductSelectAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.distributor.product.select"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelDistributorProductSelectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelDistributorProductSelectAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAscpChannelDistributorProductSelectAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpChannelDistributorProductSelectRequest()
	},
}

// GetAlibabaAscpChannelDistributorProductSelectRequest 从 sync.Pool 获取 AlibabaAscpChannelDistributorProductSelectAPIRequest
func GetAlibabaAscpChannelDistributorProductSelectAPIRequest() *AlibabaAscpChannelDistributorProductSelectAPIRequest {
	return poolAlibabaAscpChannelDistributorProductSelectAPIRequest.Get().(*AlibabaAscpChannelDistributorProductSelectAPIRequest)
}

// ReleaseAlibabaAscpChannelDistributorProductSelectAPIRequest 将 AlibabaAscpChannelDistributorProductSelectAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpChannelDistributorProductSelectAPIRequest(v *AlibabaAscpChannelDistributorProductSelectAPIRequest) {
	v.Reset()
	poolAlibabaAscpChannelDistributorProductSelectAPIRequest.Put(v)
}

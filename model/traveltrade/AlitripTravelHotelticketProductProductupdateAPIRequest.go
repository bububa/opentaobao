package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelhotelticketproductproductupdateAPIRequest 产品批量变更通知 API请求
// alitrip.travel.hotelticket.product.productupdate
//
// 产品批量变更通知
type AlitriptravelhotelticketproductproductupdateAPIRequest struct {
	model.Params
	// 系统商分配给飞猪卖家的账号
	_accessKey string
	// 变更的产品信息
	_productUpdates *ProductUpdateDto
}

// NewAlitriptravelhotelticketproductproductupdateRequest 初始化AlitriptravelhotelticketproductproductupdateAPIRequest对象
func NewAlitriptravelhotelticketproductproductupdateRequest() *AlitriptravelhotelticketproductproductupdateAPIRequest {
	return &AlitriptravelhotelticketproductproductupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptravelhotelticketproductproductupdateAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.hotelticket.product.productupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptravelhotelticketproductproductupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptravelhotelticketproductproductupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccessKey is AccessKey Setter
// 系统商分配给飞猪卖家的账号
func (r *AlitriptravelhotelticketproductproductupdateAPIRequest) SetAccessKey(_accessKey string) error {
	r._accessKey = _accessKey
	r.Set("access_key", _accessKey)
	return nil
}

// GetAccessKey AccessKey Getter
func (r AlitriptravelhotelticketproductproductupdateAPIRequest) GetAccessKey() string {
	return r._accessKey
}

// SetProductUpdates is ProductUpdates Setter
// 变更的产品信息
func (r *AlitriptravelhotelticketproductproductupdateAPIRequest) SetProductUpdates(_productUpdates *ProductUpdateDto) error {
	r._productUpdates = _productUpdates
	r.Set("product_updates", _productUpdates)
	return nil
}

// GetProductUpdates ProductUpdates Getter
func (r AlitriptravelhotelticketproductproductupdateAPIRequest) GetProductUpdates() *ProductUpdateDto {
	return r._productUpdates
}

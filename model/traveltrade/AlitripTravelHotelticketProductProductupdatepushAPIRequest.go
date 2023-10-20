package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelhotelticketproductproductupdatepushAPIRequest 产品批量变更推送通知 API请求
// alitrip.travel.hotelticket.product.productupdatepush
//
// 产品批量变更推送通知
type AlitriptravelhotelticketproductproductupdatepushAPIRequest struct {
	model.Params
	// 系统商分配给飞猪卖家的账号
	_accessKey string
	// 变更的产品信息
	_productUpdates *ProductUpdatePushDto
}

// NewAlitriptravelhotelticketproductproductupdatepushRequest 初始化AlitriptravelhotelticketproductproductupdatepushAPIRequest对象
func NewAlitriptravelhotelticketproductproductupdatepushRequest() *AlitriptravelhotelticketproductproductupdatepushAPIRequest {
	return &AlitriptravelhotelticketproductproductupdatepushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptravelhotelticketproductproductupdatepushAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.hotelticket.product.productupdatepush"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptravelhotelticketproductproductupdatepushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptravelhotelticketproductproductupdatepushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccessKey is AccessKey Setter
// 系统商分配给飞猪卖家的账号
func (r *AlitriptravelhotelticketproductproductupdatepushAPIRequest) SetAccessKey(_accessKey string) error {
	r._accessKey = _accessKey
	r.Set("access_key", _accessKey)
	return nil
}

// GetAccessKey AccessKey Getter
func (r AlitriptravelhotelticketproductproductupdatepushAPIRequest) GetAccessKey() string {
	return r._accessKey
}

// SetProductUpdates is ProductUpdates Setter
// 变更的产品信息
func (r *AlitriptravelhotelticketproductproductupdatepushAPIRequest) SetProductUpdates(_productUpdates *ProductUpdatePushDto) error {
	r._productUpdates = _productUpdates
	r.Set("product_updates", _productUpdates)
	return nil
}

// GetProductUpdates ProductUpdates Getter
func (r AlitriptravelhotelticketproductproductupdatepushAPIRequest) GetProductUpdates() *ProductUpdatePushDto {
	return r._productUpdates
}

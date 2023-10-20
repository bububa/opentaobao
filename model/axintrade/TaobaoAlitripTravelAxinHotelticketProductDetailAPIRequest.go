package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelticketproductdetailAPIRequest 阿信酒景套餐产品详情查询 API请求
// taobao.alitrip.travel.axin.hotelticket.product.detail
//
// 阿信酒景套餐产品详情查询
type TaobaoalitriptravelaxinhotelticketproductdetailAPIRequest struct {
	model.Params
	// 分销商id
	_distributorTid int64
	// 产品id
	_productId int64
}

// NewTaobaoalitriptravelaxinhotelticketproductdetailRequest 初始化TaobaoalitriptravelaxinhotelticketproductdetailAPIRequest对象
func NewTaobaoalitriptravelaxinhotelticketproductdetailRequest() *TaobaoalitriptravelaxinhotelticketproductdetailAPIRequest {
	return &TaobaoalitriptravelaxinhotelticketproductdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinhotelticketproductdetailAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotelticket.product.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinhotelticketproductdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinhotelticketproductdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoalitriptravelaxinhotelticketproductdetailAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoalitriptravelaxinhotelticketproductdetailAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetProductId is ProductId Setter
// 产品id
func (r *TaobaoalitriptravelaxinhotelticketproductdetailAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoalitriptravelaxinhotelticketproductdetailAPIRequest) GetProductId() int64 {
	return r._productId
}

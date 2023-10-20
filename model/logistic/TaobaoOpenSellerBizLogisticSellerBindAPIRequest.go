package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenSellerBizLogisticSellerBindAPIRequest 店铺授权发货注册（催发货） API请求
// taobao.open.seller.biz.logistic.seller.bind
//
// 店铺授权发货注册（催发货）
type TaobaoOpenSellerBizLogisticSellerBindAPIRequest struct {
	model.Params
	// 淘宝测试店铺Nick
	_sellerNick string
}

// NewTaobaoOpenSellerBizLogisticSellerBindRequest 初始化TaobaoOpenSellerBizLogisticSellerBindAPIRequest对象
func NewTaobaoOpenSellerBizLogisticSellerBindRequest() *TaobaoOpenSellerBizLogisticSellerBindAPIRequest {
	return &TaobaoOpenSellerBizLogisticSellerBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenSellerBizLogisticSellerBindAPIRequest) GetApiMethodName() string {
	return "taobao.open.seller.biz.logistic.seller.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenSellerBizLogisticSellerBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenSellerBizLogisticSellerBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerNick is SellerNick Setter
// 淘宝测试店铺Nick
func (r *TaobaoOpenSellerBizLogisticSellerBindAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// GetSellerNick SellerNick Getter
func (r TaobaoOpenSellerBizLogisticSellerBindAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

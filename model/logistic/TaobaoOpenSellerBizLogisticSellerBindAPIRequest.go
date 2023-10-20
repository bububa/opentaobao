package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopensellerbizlogisticsellerbindAPIRequest 店铺授权发货注册（催发货） API请求
// taobao.open.seller.biz.logistic.seller.bind
//
// 店铺授权发货注册（催发货）
type TaobaoopensellerbizlogisticsellerbindAPIRequest struct {
	model.Params
	// 淘宝测试店铺Nick
	_sellerNick string
}

// NewTaobaoopensellerbizlogisticsellerbindRequest 初始化TaobaoopensellerbizlogisticsellerbindAPIRequest对象
func NewTaobaoopensellerbizlogisticsellerbindRequest() *TaobaoopensellerbizlogisticsellerbindAPIRequest {
	return &TaobaoopensellerbizlogisticsellerbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopensellerbizlogisticsellerbindAPIRequest) GetApiMethodName() string {
	return "taobao.open.seller.biz.logistic.seller.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopensellerbizlogisticsellerbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopensellerbizlogisticsellerbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerNick is SellerNick Setter
// 淘宝测试店铺Nick
func (r *TaobaoopensellerbizlogisticsellerbindAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// GetSellerNick SellerNick Getter
func (r TaobaoopensellerbizlogisticsellerbindAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

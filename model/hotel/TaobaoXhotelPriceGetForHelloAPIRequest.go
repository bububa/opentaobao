package hotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelpricegetforhelloAPIRequest 哈罗合作方获取酒店库存报价 API请求
// taobao.xhotel.price.get.for.hello
//
// 哈罗合作项目，供哈罗合作方按需查询已开通城市下的标准酒店下指定时间段内的库存报价信息；在用户登录方面，返回结果不涉及用户个人信息，不涉及商家信息；仅根据不同用户，查询对应会员等级后，返回不同报价；
type TaobaoxhotelpricegetforhelloAPIRequest struct {
	model.Params
	// 参数封装
	_hotelPriceParam *HotelPriceParam
}

// NewTaobaoxhotelpricegetforhelloRequest 初始化TaobaoxhotelpricegetforhelloAPIRequest对象
func NewTaobaoxhotelpricegetforhelloRequest() *TaobaoxhotelpricegetforhelloAPIRequest {
	return &TaobaoxhotelpricegetforhelloAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelpricegetforhelloAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.price.get.for.hello"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelpricegetforhelloAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelpricegetforhelloAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHotelPriceParam is HotelPriceParam Setter
// 参数封装
func (r *TaobaoxhotelpricegetforhelloAPIRequest) SetHotelPriceParam(_hotelPriceParam *HotelPriceParam) error {
	r._hotelPriceParam = _hotelPriceParam
	r.Set("hotel_price_param", _hotelPriceParam)
	return nil
}

// GetHotelPriceParam HotelPriceParam Getter
func (r TaobaoxhotelpricegetforhelloAPIRequest) GetHotelPriceParam() *HotelPriceParam {
	return r._hotelPriceParam
}

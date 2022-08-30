package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelMatchAPIRequest 酒店匹配接口-阿信 API请求
// taobao.alitrip.travel.axin.hotel.match
//
// 酒店匹配接口-阿信
type TaobaoAlitripTravelAxinHotelMatchAPIRequest struct {
	model.Params
	// 请求
	_req *MatchedHotelRequestDto
}

// NewTaobaoAlitripTravelAxinHotelMatchRequest 初始化TaobaoAlitripTravelAxinHotelMatchAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelMatchRequest() *TaobaoAlitripTravelAxinHotelMatchAPIRequest {
	return &TaobaoAlitripTravelAxinHotelMatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelMatchAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.match"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelMatchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReq is Req Setter
// 请求
func (r *TaobaoAlitripTravelAxinHotelMatchAPIRequest) SetReq(_req *MatchedHotelRequestDto) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r TaobaoAlitripTravelAxinHotelMatchAPIRequest) GetReq() *MatchedHotelRequestDto {
	return r._req
}

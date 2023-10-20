package axindata

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelAxinHotelMatchAPIRequest) Reset() {
	r._req = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelMatchAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.match"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelMatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelAxinHotelMatchAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoAlitripTravelAxinHotelMatchAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelAxinHotelMatchRequest()
	},
}

// GetTaobaoAlitripTravelAxinHotelMatchRequest 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelMatchAPIRequest
func GetTaobaoAlitripTravelAxinHotelMatchAPIRequest() *TaobaoAlitripTravelAxinHotelMatchAPIRequest {
	return poolTaobaoAlitripTravelAxinHotelMatchAPIRequest.Get().(*TaobaoAlitripTravelAxinHotelMatchAPIRequest)
}

// ReleaseTaobaoAlitripTravelAxinHotelMatchAPIRequest 将 TaobaoAlitripTravelAxinHotelMatchAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelMatchAPIRequest(v *TaobaoAlitripTravelAxinHotelMatchAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelMatchAPIRequest.Put(v)
}

package hotelhstdf

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfShotelMatchshotelselfAPIRequest 自主匹配标准酒店以及卖家酒店 API请求
// alitrip.hotel.hstdf.shotel.matchshotelself
//
// 商家通过指定的标准酒店id和卖家酒店id进行匹配
type AlitripHotelHstdfShotelMatchshotelselfAPIRequest struct {
	model.Params
	// HotelMatchParam
	_param0 *HotelMatchParam
}

// NewAlitripHotelHstdfShotelMatchshotelselfRequest 初始化AlitripHotelHstdfShotelMatchshotelselfAPIRequest对象
func NewAlitripHotelHstdfShotelMatchshotelselfRequest() *AlitripHotelHstdfShotelMatchshotelselfAPIRequest {
	return &AlitripHotelHstdfShotelMatchshotelselfAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripHotelHstdfShotelMatchshotelselfAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfShotelMatchshotelselfAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.shotel.matchshotelself"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfShotelMatchshotelselfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripHotelHstdfShotelMatchshotelselfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// HotelMatchParam
func (r *AlitripHotelHstdfShotelMatchshotelselfAPIRequest) SetParam0(_param0 *HotelMatchParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlitripHotelHstdfShotelMatchshotelselfAPIRequest) GetParam0() *HotelMatchParam {
	return r._param0
}

var poolAlitripHotelHstdfShotelMatchshotelselfAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripHotelHstdfShotelMatchshotelselfRequest()
	},
}

// GetAlitripHotelHstdfShotelMatchshotelselfRequest 从 sync.Pool 获取 AlitripHotelHstdfShotelMatchshotelselfAPIRequest
func GetAlitripHotelHstdfShotelMatchshotelselfAPIRequest() *AlitripHotelHstdfShotelMatchshotelselfAPIRequest {
	return poolAlitripHotelHstdfShotelMatchshotelselfAPIRequest.Get().(*AlitripHotelHstdfShotelMatchshotelselfAPIRequest)
}

// ReleaseAlitripHotelHstdfShotelMatchshotelselfAPIRequest 将 AlitripHotelHstdfShotelMatchshotelselfAPIRequest 放入 sync.Pool
func ReleaseAlitripHotelHstdfShotelMatchshotelselfAPIRequest(v *AlitripHotelHstdfShotelMatchshotelselfAPIRequest) {
	v.Reset()
	poolAlitripHotelHstdfShotelMatchshotelselfAPIRequest.Put(v)
}

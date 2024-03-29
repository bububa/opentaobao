package hotel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelInfoListGetForHelloAPIRequest 哈罗获取酒店详情 API请求
// taobao.xhotel.info.list.get.for.hello
//
// 哈罗合作项目，供哈罗合作方批量和增量两种场景下查询已开通城市下的标准酒店及房型信息，不涉及用户登录信息
type TaobaoXhotelInfoListGetForHelloAPIRequest struct {
	model.Params
	// 参数封装模型
	_hotelInfoParam *HotelInfoParam
}

// NewTaobaoXhotelInfoListGetForHelloRequest 初始化TaobaoXhotelInfoListGetForHelloAPIRequest对象
func NewTaobaoXhotelInfoListGetForHelloRequest() *TaobaoXhotelInfoListGetForHelloAPIRequest {
	return &TaobaoXhotelInfoListGetForHelloAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelInfoListGetForHelloAPIRequest) Reset() {
	r._hotelInfoParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelInfoListGetForHelloAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.info.list.get.for.hello"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelInfoListGetForHelloAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelInfoListGetForHelloAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHotelInfoParam is HotelInfoParam Setter
// 参数封装模型
func (r *TaobaoXhotelInfoListGetForHelloAPIRequest) SetHotelInfoParam(_hotelInfoParam *HotelInfoParam) error {
	r._hotelInfoParam = _hotelInfoParam
	r.Set("hotel_info_param", _hotelInfoParam)
	return nil
}

// GetHotelInfoParam HotelInfoParam Getter
func (r TaobaoXhotelInfoListGetForHelloAPIRequest) GetHotelInfoParam() *HotelInfoParam {
	return r._hotelInfoParam
}

var poolTaobaoXhotelInfoListGetForHelloAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelInfoListGetForHelloRequest()
	},
}

// GetTaobaoXhotelInfoListGetForHelloRequest 从 sync.Pool 获取 TaobaoXhotelInfoListGetForHelloAPIRequest
func GetTaobaoXhotelInfoListGetForHelloAPIRequest() *TaobaoXhotelInfoListGetForHelloAPIRequest {
	return poolTaobaoXhotelInfoListGetForHelloAPIRequest.Get().(*TaobaoXhotelInfoListGetForHelloAPIRequest)
}

// ReleaseTaobaoXhotelInfoListGetForHelloAPIRequest 将 TaobaoXhotelInfoListGetForHelloAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelInfoListGetForHelloAPIRequest(v *TaobaoXhotelInfoListGetForHelloAPIRequest) {
	v.Reset()
	poolTaobaoXhotelInfoListGetForHelloAPIRequest.Put(v)
}

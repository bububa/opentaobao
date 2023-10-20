package hotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelinfolistgetforhelloAPIRequest 哈罗获取酒店详情 API请求
// taobao.xhotel.info.list.get.for.hello
//
// 哈罗合作项目，供哈罗合作方批量和增量两种场景下查询已开通城市下的标准酒店及房型信息，不涉及用户登录信息
type TaobaoxhotelinfolistgetforhelloAPIRequest struct {
	model.Params
	// 参数封装模型
	_hotelInfoParam *HotelInfoParam
}

// NewTaobaoxhotelinfolistgetforhelloRequest 初始化TaobaoxhotelinfolistgetforhelloAPIRequest对象
func NewTaobaoxhotelinfolistgetforhelloRequest() *TaobaoxhotelinfolistgetforhelloAPIRequest {
	return &TaobaoxhotelinfolistgetforhelloAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelinfolistgetforhelloAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.info.list.get.for.hello"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelinfolistgetforhelloAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelinfolistgetforhelloAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHotelInfoParam is HotelInfoParam Setter
// 参数封装模型
func (r *TaobaoxhotelinfolistgetforhelloAPIRequest) SetHotelInfoParam(_hotelInfoParam *HotelInfoParam) error {
	r._hotelInfoParam = _hotelInfoParam
	r.Set("hotel_info_param", _hotelInfoParam)
	return nil
}

// GetHotelInfoParam HotelInfoParam Getter
func (r TaobaoxhotelinfolistgetforhelloAPIRequest) GetHotelInfoParam() *HotelInfoParam {
	return r._hotelInfoParam
}

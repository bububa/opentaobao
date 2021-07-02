package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizLightUpAPIRequest 价签LED等点亮 API请求
// taobao.uscesl.biz.light.up
//
// 价签LED等点亮
type TaobaoUsceslBizLightUpAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 商家编号
	_bizBrandKey string
	// 价签条码
	_eslBarCode string
	// 亮灯颜色，绿：值为2；红：值为4
	_ledColor string
	// 亮灯时长，单位：秒，最大长度3600秒
	_lightUpTime int64
}

// NewTaobaoUsceslBizLightUpRequest 初始化TaobaoUsceslBizLightUpAPIRequest对象
func NewTaobaoUsceslBizLightUpRequest() *TaobaoUsceslBizLightUpAPIRequest {
	return &TaobaoUsceslBizLightUpAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizLightUpAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.light.up"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizLightUpAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreId Setter
// 门店ID
func (r *TaobaoUsceslBizLightUpAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r TaobaoUsceslBizLightUpAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// Set is BizBrandKey Setter
// 商家编号
func (r *TaobaoUsceslBizLightUpAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// Get BizBrandKey Getter
func (r TaobaoUsceslBizLightUpAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

// Set is EslBarCode Setter
// 价签条码
func (r *TaobaoUsceslBizLightUpAPIRequest) SetEslBarCode(_eslBarCode string) error {
	r._eslBarCode = _eslBarCode
	r.Set("esl_bar_code", _eslBarCode)
	return nil
}

// Get EslBarCode Getter
func (r TaobaoUsceslBizLightUpAPIRequest) GetEslBarCode() string {
	return r._eslBarCode
}

// Set is LedColor Setter
// 亮灯颜色，绿：值为2；红：值为4
func (r *TaobaoUsceslBizLightUpAPIRequest) SetLedColor(_ledColor string) error {
	r._ledColor = _ledColor
	r.Set("led_color", _ledColor)
	return nil
}

// Get LedColor Getter
func (r TaobaoUsceslBizLightUpAPIRequest) GetLedColor() string {
	return r._ledColor
}

// Set is LightUpTime Setter
// 亮灯时长，单位：秒，最大长度3600秒
func (r *TaobaoUsceslBizLightUpAPIRequest) SetLightUpTime(_lightUpTime int64) error {
	r._lightUpTime = _lightUpTime
	r.Set("light_up_time", _lightUpTime)
	return nil
}

// Get LightUpTime Getter
func (r TaobaoUsceslBizLightUpAPIRequest) GetLightUpTime() int64 {
	return r._lightUpTime
}

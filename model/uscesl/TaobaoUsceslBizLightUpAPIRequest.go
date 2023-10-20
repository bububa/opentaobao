package uscesl

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizLightUpAPIRequest 价签LED等点亮 API请求
// taobao.uscesl.biz.light.up
//
// 价签LED等点亮
type TaobaoUsceslBizLightUpAPIRequest struct {
	model.Params
	// 价签条码
	_eslBarCode string
	// 亮灯颜色，绿：值为2；红：值为4
	_ledColor string
	// 商家编号
	_bizBrandKey string
	// 亮灯时长，单位：秒，最大长度3600秒
	_lightUpTime int64
	// 门店ID
	_storeId int64
}

// NewTaobaoUsceslBizLightUpRequest 初始化TaobaoUsceslBizLightUpAPIRequest对象
func NewTaobaoUsceslBizLightUpRequest() *TaobaoUsceslBizLightUpAPIRequest {
	return &TaobaoUsceslBizLightUpAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUsceslBizLightUpAPIRequest) Reset() {
	r._eslBarCode = ""
	r._ledColor = ""
	r._bizBrandKey = ""
	r._lightUpTime = 0
	r._storeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizLightUpAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.light.up"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizLightUpAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUsceslBizLightUpAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEslBarCode is EslBarCode Setter
// 价签条码
func (r *TaobaoUsceslBizLightUpAPIRequest) SetEslBarCode(_eslBarCode string) error {
	r._eslBarCode = _eslBarCode
	r.Set("esl_bar_code", _eslBarCode)
	return nil
}

// GetEslBarCode EslBarCode Getter
func (r TaobaoUsceslBizLightUpAPIRequest) GetEslBarCode() string {
	return r._eslBarCode
}

// SetLedColor is LedColor Setter
// 亮灯颜色，绿：值为2；红：值为4
func (r *TaobaoUsceslBizLightUpAPIRequest) SetLedColor(_ledColor string) error {
	r._ledColor = _ledColor
	r.Set("led_color", _ledColor)
	return nil
}

// GetLedColor LedColor Getter
func (r TaobaoUsceslBizLightUpAPIRequest) GetLedColor() string {
	return r._ledColor
}

// SetBizBrandKey is BizBrandKey Setter
// 商家编号
func (r *TaobaoUsceslBizLightUpAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// GetBizBrandKey BizBrandKey Getter
func (r TaobaoUsceslBizLightUpAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

// SetLightUpTime is LightUpTime Setter
// 亮灯时长，单位：秒，最大长度3600秒
func (r *TaobaoUsceslBizLightUpAPIRequest) SetLightUpTime(_lightUpTime int64) error {
	r._lightUpTime = _lightUpTime
	r.Set("light_up_time", _lightUpTime)
	return nil
}

// GetLightUpTime LightUpTime Getter
func (r TaobaoUsceslBizLightUpAPIRequest) GetLightUpTime() int64 {
	return r._lightUpTime
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoUsceslBizLightUpAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoUsceslBizLightUpAPIRequest) GetStoreId() int64 {
	return r._storeId
}

var poolTaobaoUsceslBizLightUpAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUsceslBizLightUpRequest()
	},
}

// GetTaobaoUsceslBizLightUpRequest 从 sync.Pool 获取 TaobaoUsceslBizLightUpAPIRequest
func GetTaobaoUsceslBizLightUpAPIRequest() *TaobaoUsceslBizLightUpAPIRequest {
	return poolTaobaoUsceslBizLightUpAPIRequest.Get().(*TaobaoUsceslBizLightUpAPIRequest)
}

// ReleaseTaobaoUsceslBizLightUpAPIRequest 将 TaobaoUsceslBizLightUpAPIRequest 放入 sync.Pool
func ReleaseTaobaoUsceslBizLightUpAPIRequest(v *TaobaoUsceslBizLightUpAPIRequest) {
	v.Reset()
	poolTaobaoUsceslBizLightUpAPIRequest.Put(v)
}

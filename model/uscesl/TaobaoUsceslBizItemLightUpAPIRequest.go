package uscesl

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizItemLightUpAPIRequest 商品条码亮灯API API请求
// taobao.uscesl.biz.item.light.up
//
// 亮灯API
type TaobaoUsceslBizItemLightUpAPIRequest struct {
	model.Params
	// 商品条码
	_itemBarCode string
	// 亮灯颜色，绿：值为2；红：值为4
	_ledColor string
	// 商家编号
	_bizBrandKey string
	// 亮灯时长，单位：秒，最大长度3600秒
	_lightUpTime int64
	// 门店编号
	_storeId int64
}

// NewTaobaoUsceslBizItemLightUpRequest 初始化TaobaoUsceslBizItemLightUpAPIRequest对象
func NewTaobaoUsceslBizItemLightUpRequest() *TaobaoUsceslBizItemLightUpAPIRequest {
	return &TaobaoUsceslBizItemLightUpAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUsceslBizItemLightUpAPIRequest) Reset() {
	r._itemBarCode = ""
	r._ledColor = ""
	r._bizBrandKey = ""
	r._lightUpTime = 0
	r._storeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizItemLightUpAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.item.light.up"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizItemLightUpAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUsceslBizItemLightUpAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemBarCode is ItemBarCode Setter
// 商品条码
func (r *TaobaoUsceslBizItemLightUpAPIRequest) SetItemBarCode(_itemBarCode string) error {
	r._itemBarCode = _itemBarCode
	r.Set("item_bar_code", _itemBarCode)
	return nil
}

// GetItemBarCode ItemBarCode Getter
func (r TaobaoUsceslBizItemLightUpAPIRequest) GetItemBarCode() string {
	return r._itemBarCode
}

// SetLedColor is LedColor Setter
// 亮灯颜色，绿：值为2；红：值为4
func (r *TaobaoUsceslBizItemLightUpAPIRequest) SetLedColor(_ledColor string) error {
	r._ledColor = _ledColor
	r.Set("led_color", _ledColor)
	return nil
}

// GetLedColor LedColor Getter
func (r TaobaoUsceslBizItemLightUpAPIRequest) GetLedColor() string {
	return r._ledColor
}

// SetBizBrandKey is BizBrandKey Setter
// 商家编号
func (r *TaobaoUsceslBizItemLightUpAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// GetBizBrandKey BizBrandKey Getter
func (r TaobaoUsceslBizItemLightUpAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

// SetLightUpTime is LightUpTime Setter
// 亮灯时长，单位：秒，最大长度3600秒
func (r *TaobaoUsceslBizItemLightUpAPIRequest) SetLightUpTime(_lightUpTime int64) error {
	r._lightUpTime = _lightUpTime
	r.Set("light_up_time", _lightUpTime)
	return nil
}

// GetLightUpTime LightUpTime Getter
func (r TaobaoUsceslBizItemLightUpAPIRequest) GetLightUpTime() int64 {
	return r._lightUpTime
}

// SetStoreId is StoreId Setter
// 门店编号
func (r *TaobaoUsceslBizItemLightUpAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoUsceslBizItemLightUpAPIRequest) GetStoreId() int64 {
	return r._storeId
}

var poolTaobaoUsceslBizItemLightUpAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUsceslBizItemLightUpRequest()
	},
}

// GetTaobaoUsceslBizItemLightUpRequest 从 sync.Pool 获取 TaobaoUsceslBizItemLightUpAPIRequest
func GetTaobaoUsceslBizItemLightUpAPIRequest() *TaobaoUsceslBizItemLightUpAPIRequest {
	return poolTaobaoUsceslBizItemLightUpAPIRequest.Get().(*TaobaoUsceslBizItemLightUpAPIRequest)
}

// ReleaseTaobaoUsceslBizItemLightUpAPIRequest 将 TaobaoUsceslBizItemLightUpAPIRequest 放入 sync.Pool
func ReleaseTaobaoUsceslBizItemLightUpAPIRequest(v *TaobaoUsceslBizItemLightUpAPIRequest) {
	v.Reset()
	poolTaobaoUsceslBizItemLightUpAPIRequest.Put(v)
}

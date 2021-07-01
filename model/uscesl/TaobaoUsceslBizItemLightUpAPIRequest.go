package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsceslBizItemLightUpAPIRequest
商品条码亮灯API API请求
taobao.uscesl.biz.item.light.up

亮灯API */
type TaobaoUsceslBizItemLightUpAPIRequest struct {
	model.Params
	// 商品条码
	_itemBarCode string
	// 亮灯颜色，绿：值为2；红：值为4
	_ledColor string
	// 亮灯时长，单位：秒，最大长度3600秒
	_lightUpTime int64
	// 门店编号
	_storeId int64
	// 商家编号
	_bizBrandKey string
}

// NewTaobaoUsceslBizItemLightUpRequest 初始化TaobaoUsceslBizItemLightUpAPIRequest对象
func NewTaobaoUsceslBizItemLightUpRequest() *TaobaoUsceslBizItemLightUpAPIRequest {
	return &TaobaoUsceslBizItemLightUpAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizItemLightUpAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.item.light.up"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizItemLightUpAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemBarCode Setter
// 商品条码
func (r *TaobaoUsceslBizItemLightUpAPIRequest) SetItemBarCode(_itemBarCode string) error {
	r._itemBarCode = _itemBarCode
	r.Set("item_bar_code", _itemBarCode)
	return nil
}

// Get ItemBarCode Getter
func (r TaobaoUsceslBizItemLightUpAPIRequest) GetItemBarCode() string {
	return r._itemBarCode
}

// Set is LedColor Setter
// 亮灯颜色，绿：值为2；红：值为4
func (r *TaobaoUsceslBizItemLightUpAPIRequest) SetLedColor(_ledColor string) error {
	r._ledColor = _ledColor
	r.Set("led_color", _ledColor)
	return nil
}

// Get LedColor Getter
func (r TaobaoUsceslBizItemLightUpAPIRequest) GetLedColor() string {
	return r._ledColor
}

// Set is LightUpTime Setter
// 亮灯时长，单位：秒，最大长度3600秒
func (r *TaobaoUsceslBizItemLightUpAPIRequest) SetLightUpTime(_lightUpTime int64) error {
	r._lightUpTime = _lightUpTime
	r.Set("light_up_time", _lightUpTime)
	return nil
}

// Get LightUpTime Getter
func (r TaobaoUsceslBizItemLightUpAPIRequest) GetLightUpTime() int64 {
	return r._lightUpTime
}

// Set is StoreId Setter
// 门店编号
func (r *TaobaoUsceslBizItemLightUpAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r TaobaoUsceslBizItemLightUpAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// Set is BizBrandKey Setter
// 商家编号
func (r *TaobaoUsceslBizItemLightUpAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// Get BizBrandKey Getter
func (r TaobaoUsceslBizItemLightUpAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

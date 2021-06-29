package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品条码亮灯API API请求
taobao.uscesl.biz.item.light.up

亮灯API
*/
type TaobaoUsceslBizItemLightUpRequest struct {
    model.Params
    // 商品条码
    _itemBarCode   string
    // 亮灯颜色，绿：值为2；红：值为4
    _ledColor   string
    // 亮灯时长，单位：秒，最大长度3600秒
    _lightUpTime   int64
    // 门店编号
    _storeId   int64
    // 商家编号
    _bizBrandKey   string
}

// 初始化TaobaoUsceslBizItemLightUpRequest对象
func NewTaobaoUsceslBizItemLightUpRequest() *TaobaoUsceslBizItemLightUpRequest{
    return &TaobaoUsceslBizItemLightUpRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizItemLightUpRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.item.light.up"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizItemLightUpRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemBarCode Setter
// 商品条码
func (r *TaobaoUsceslBizItemLightUpRequest) SetItemBarCode(_itemBarCode string) error {
    r._itemBarCode = _itemBarCode
    r.Set("item_bar_code", _itemBarCode)
    return nil
}

// ItemBarCode Getter
func (r TaobaoUsceslBizItemLightUpRequest) GetItemBarCode() string {
    return r._itemBarCode
}
// LedColor Setter
// 亮灯颜色，绿：值为2；红：值为4
func (r *TaobaoUsceslBizItemLightUpRequest) SetLedColor(_ledColor string) error {
    r._ledColor = _ledColor
    r.Set("led_color", _ledColor)
    return nil
}

// LedColor Getter
func (r TaobaoUsceslBizItemLightUpRequest) GetLedColor() string {
    return r._ledColor
}
// LightUpTime Setter
// 亮灯时长，单位：秒，最大长度3600秒
func (r *TaobaoUsceslBizItemLightUpRequest) SetLightUpTime(_lightUpTime int64) error {
    r._lightUpTime = _lightUpTime
    r.Set("light_up_time", _lightUpTime)
    return nil
}

// LightUpTime Getter
func (r TaobaoUsceslBizItemLightUpRequest) GetLightUpTime() int64 {
    return r._lightUpTime
}
// StoreId Setter
// 门店编号
func (r *TaobaoUsceslBizItemLightUpRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoUsceslBizItemLightUpRequest) GetStoreId() int64 {
    return r._storeId
}
// BizBrandKey Setter
// 商家编号
func (r *TaobaoUsceslBizItemLightUpRequest) SetBizBrandKey(_bizBrandKey string) error {
    r._bizBrandKey = _bizBrandKey
    r.Set("biz_brand_key", _bizBrandKey)
    return nil
}

// BizBrandKey Getter
func (r TaobaoUsceslBizItemLightUpRequest) GetBizBrandKey() string {
    return r._bizBrandKey
}

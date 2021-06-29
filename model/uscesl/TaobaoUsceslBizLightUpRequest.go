package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
价签LED等点亮 API请求
taobao.uscesl.biz.light.up

价签LED等点亮
*/
type TaobaoUsceslBizLightUpRequest struct {
    model.Params
    // 门店ID
    storeId   int64
    // 商家编号
    bizBrandKey   string
    // 价签条码
    eslBarCode   string
    // 亮灯颜色，绿：值为2；红：值为4
    ledColor   string
    // 亮灯时长，单位：秒，最大长度3600秒
    lightUpTime   int64
}

// 初始化TaobaoUsceslBizLightUpRequest对象
func NewTaobaoUsceslBizLightUpRequest() *TaobaoUsceslBizLightUpRequest{
    return &TaobaoUsceslBizLightUpRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizLightUpRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.light.up"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizLightUpRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店ID
func (r *TaobaoUsceslBizLightUpRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoUsceslBizLightUpRequest) GetStoreId() int64 {
    return r.storeId
}
// BizBrandKey Setter
// 商家编号
func (r *TaobaoUsceslBizLightUpRequest) SetBizBrandKey(bizBrandKey string) error {
    r.bizBrandKey = bizBrandKey
    r.Set("biz_brand_key", bizBrandKey)
    return nil
}

// BizBrandKey Getter
func (r TaobaoUsceslBizLightUpRequest) GetBizBrandKey() string {
    return r.bizBrandKey
}
// EslBarCode Setter
// 价签条码
func (r *TaobaoUsceslBizLightUpRequest) SetEslBarCode(eslBarCode string) error {
    r.eslBarCode = eslBarCode
    r.Set("esl_bar_code", eslBarCode)
    return nil
}

// EslBarCode Getter
func (r TaobaoUsceslBizLightUpRequest) GetEslBarCode() string {
    return r.eslBarCode
}
// LedColor Setter
// 亮灯颜色，绿：值为2；红：值为4
func (r *TaobaoUsceslBizLightUpRequest) SetLedColor(ledColor string) error {
    r.ledColor = ledColor
    r.Set("led_color", ledColor)
    return nil
}

// LedColor Getter
func (r TaobaoUsceslBizLightUpRequest) GetLedColor() string {
    return r.ledColor
}
// LightUpTime Setter
// 亮灯时长，单位：秒，最大长度3600秒
func (r *TaobaoUsceslBizLightUpRequest) SetLightUpTime(lightUpTime int64) error {
    r.lightUpTime = lightUpTime
    r.Set("light_up_time", lightUpTime)
    return nil
}

// LightUpTime Getter
func (r TaobaoUsceslBizLightUpRequest) GetLightUpTime() int64 {
    return r.lightUpTime
}

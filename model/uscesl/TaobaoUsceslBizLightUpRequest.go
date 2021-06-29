package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
价签LED等点亮 APIRequest
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

func NewTaobaoUsceslBizLightUpRequest() *TaobaoUsceslBizLightUpRequest{
    return &TaobaoUsceslBizLightUpRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsceslBizLightUpRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.light.up"
}

func (r TaobaoUsceslBizLightUpRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsceslBizLightUpRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoUsceslBizLightUpRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoUsceslBizLightUpRequest) SetBizBrandKey(bizBrandKey string) error {
    r.bizBrandKey = bizBrandKey
    r.Set("biz_brand_key", bizBrandKey)
    return nil
}

func (r TaobaoUsceslBizLightUpRequest) GetBizBrandKey() string {
    return r.bizBrandKey
}

func (r *TaobaoUsceslBizLightUpRequest) SetEslBarCode(eslBarCode string) error {
    r.eslBarCode = eslBarCode
    r.Set("esl_bar_code", eslBarCode)
    return nil
}

func (r TaobaoUsceslBizLightUpRequest) GetEslBarCode() string {
    return r.eslBarCode
}

func (r *TaobaoUsceslBizLightUpRequest) SetLedColor(ledColor string) error {
    r.ledColor = ledColor
    r.Set("led_color", ledColor)
    return nil
}

func (r TaobaoUsceslBizLightUpRequest) GetLedColor() string {
    return r.ledColor
}

func (r *TaobaoUsceslBizLightUpRequest) SetLightUpTime(lightUpTime int64) error {
    r.lightUpTime = lightUpTime
    r.Set("light_up_time", lightUpTime)
    return nil
}

func (r TaobaoUsceslBizLightUpRequest) GetLightUpTime() int64 {
    return r.lightUpTime
}


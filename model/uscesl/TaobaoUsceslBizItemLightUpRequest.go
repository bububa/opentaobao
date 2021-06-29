package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品条码亮灯API APIRequest
taobao.uscesl.biz.item.light.up

亮灯API
*/
type TaobaoUsceslBizItemLightUpRequest struct {
    model.Params

    // 商品条码
    itemBarCode   string 

    // 亮灯颜色，绿：值为2；红：值为4
    ledColor   string 

    // 亮灯时长，单位：秒，最大长度3600秒
    lightUpTime   int64 

    // 门店编号
    storeId   int64 

    // 商家编号
    bizBrandKey   string 

}

func NewTaobaoUsceslBizItemLightUpRequest() *TaobaoUsceslBizItemLightUpRequest{
    return &TaobaoUsceslBizItemLightUpRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsceslBizItemLightUpRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.item.light.up"
}

func (r TaobaoUsceslBizItemLightUpRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsceslBizItemLightUpRequest) SetItemBarCode(itemBarCode string) error {
    r.itemBarCode = itemBarCode
    r.Set("item_bar_code", itemBarCode)
    return nil
}

func (r TaobaoUsceslBizItemLightUpRequest) GetItemBarCode() string {
    return r.itemBarCode
}

func (r *TaobaoUsceslBizItemLightUpRequest) SetLedColor(ledColor string) error {
    r.ledColor = ledColor
    r.Set("led_color", ledColor)
    return nil
}

func (r TaobaoUsceslBizItemLightUpRequest) GetLedColor() string {
    return r.ledColor
}

func (r *TaobaoUsceslBizItemLightUpRequest) SetLightUpTime(lightUpTime int64) error {
    r.lightUpTime = lightUpTime
    r.Set("light_up_time", lightUpTime)
    return nil
}

func (r TaobaoUsceslBizItemLightUpRequest) GetLightUpTime() int64 {
    return r.lightUpTime
}

func (r *TaobaoUsceslBizItemLightUpRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoUsceslBizItemLightUpRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoUsceslBizItemLightUpRequest) SetBizBrandKey(bizBrandKey string) error {
    r.bizBrandKey = bizBrandKey
    r.Set("biz_brand_key", bizBrandKey)
    return nil
}

func (r TaobaoUsceslBizItemLightUpRequest) GetBizBrandKey() string {
    return r.bizBrandKey
}


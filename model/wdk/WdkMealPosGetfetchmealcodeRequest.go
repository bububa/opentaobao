package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口餐饮取餐号获取接口 APIRequest
wdk.meal.pos.getfetchmealcode

pos机创建订单前获取餐饮取餐号
*/
type WdkMealPosGetfetchmealcodeRequest struct {
    model.Params

    // 渠道店id
    channelShopId   string 

}

func NewWdkMealPosGetfetchmealcodeRequest() *WdkMealPosGetfetchmealcodeRequest{
    return &WdkMealPosGetfetchmealcodeRequest{
        Params: model.NewParams(),
    }
}

func (r WdkMealPosGetfetchmealcodeRequest) GetApiMethodName() string {
    return "wdk.meal.pos.getfetchmealcode"
}

func (r WdkMealPosGetfetchmealcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *WdkMealPosGetfetchmealcodeRequest) SetChannelShopId(channelShopId string) error {
    r.channelShopId = channelShopId
    r.Set("channel_shop_id", channelShopId)
    return nil
}

func (r WdkMealPosGetfetchmealcodeRequest) GetChannelShopId() string {
    return r.channelShopId
}


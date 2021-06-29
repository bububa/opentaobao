package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口餐饮取餐号获取接口 API请求
wdk.meal.pos.getfetchmealcode

pos机创建订单前获取餐饮取餐号
*/
type WdkMealPosGetfetchmealcodeRequest struct {
    model.Params
    // 渠道店id
    _channelShopId   string
}

// 初始化WdkMealPosGetfetchmealcodeRequest对象
func NewWdkMealPosGetfetchmealcodeRequest() *WdkMealPosGetfetchmealcodeRequest{
    return &WdkMealPosGetfetchmealcodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r WdkMealPosGetfetchmealcodeRequest) GetApiMethodName() string {
    return "wdk.meal.pos.getfetchmealcode"
}

// IRequest interface 方法, 获取API参数
func (r WdkMealPosGetfetchmealcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChannelShopId Setter
// 渠道店id
func (r *WdkMealPosGetfetchmealcodeRequest) SetChannelShopId(_channelShopId string) error {
    r._channelShopId = _channelShopId
    r.Set("channel_shop_id", _channelShopId)
    return nil
}

// ChannelShopId Getter
func (r WdkMealPosGetfetchmealcodeRequest) GetChannelShopId() string {
    return r._channelShopId
}

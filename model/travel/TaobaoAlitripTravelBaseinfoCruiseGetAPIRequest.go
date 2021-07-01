package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假线路商品发布时基础信息获取接口：邮轮扩展信息获取 API请求
taobao.alitrip.travel.baseinfo.cruise.get

旅行度假新商品发布时可用的扩展接口，用于获取邮轮类目相关扩展信息。
*/
type TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest struct {
    model.Params
    // true-获取国际邮轮类目扩展信息；false-获取国内邮轮类目扩展信息
    _isOverseas   bool
}

// 初始化TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest对象
func NewTaobaoAlitripTravelBaseinfoCruiseGetRequest() *TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest{
    return &TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.baseinfo.cruise.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IsOverseas Setter
// true-获取国际邮轮类目扩展信息；false-获取国内邮轮类目扩展信息
func (r *TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest) SetIsOverseas(_isOverseas bool) error {
    r._isOverseas = _isOverseas
    r.Set("is_overseas", _isOverseas)
    return nil
}

// IsOverseas Getter
func (r TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest) GetIsOverseas() bool {
    return r._isOverseas
}

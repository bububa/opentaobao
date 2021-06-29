package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词出价指导工具（新） API请求
taobao.simba.bidword.pricetools

关键词出价指导工具（新）
*/
type TaobaoSimbaBidwordPricetoolsRequest struct {
    model.Params
    // 关键词id
    _bidwordId   int64
    // 出价目标 ，1：争取排名；2：提升展现；3：提示点击；4：提升转化
    _type   int64
    // 区分渠道 ，计算机：PC，无线 ：WL
    _trafficType   string
    // 推广单元id
    _adgroupId   int64
}

// 初始化TaobaoSimbaBidwordPricetoolsRequest对象
func NewTaobaoSimbaBidwordPricetoolsRequest() *TaobaoSimbaBidwordPricetoolsRequest{
    return &TaobaoSimbaBidwordPricetoolsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaBidwordPricetoolsRequest) GetApiMethodName() string {
    return "taobao.simba.bidword.pricetools"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaBidwordPricetoolsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BidwordId Setter
// 关键词id
func (r *TaobaoSimbaBidwordPricetoolsRequest) SetBidwordId(_bidwordId int64) error {
    r._bidwordId = _bidwordId
    r.Set("bidword_id", _bidwordId)
    return nil
}

// BidwordId Getter
func (r TaobaoSimbaBidwordPricetoolsRequest) GetBidwordId() int64 {
    return r._bidwordId
}
// Type Setter
// 出价目标 ，1：争取排名；2：提升展现；3：提示点击；4：提升转化
func (r *TaobaoSimbaBidwordPricetoolsRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoSimbaBidwordPricetoolsRequest) GetType() int64 {
    return r._type
}
// TrafficType Setter
// 区分渠道 ，计算机：PC，无线 ：WL
func (r *TaobaoSimbaBidwordPricetoolsRequest) SetTrafficType(_trafficType string) error {
    r._trafficType = _trafficType
    r.Set("traffic_type", _trafficType)
    return nil
}

// TrafficType Getter
func (r TaobaoSimbaBidwordPricetoolsRequest) GetTrafficType() string {
    return r._trafficType
}
// AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaBidwordPricetoolsRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaBidwordPricetoolsRequest) GetAdgroupId() int64 {
    return r._adgroupId
}

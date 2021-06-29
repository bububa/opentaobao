package cityretail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
完美履约订单物流状态查询接口 API请求
tmall.cityretail.wmfl.order.logistics.query

完美履约订单物流状态查询接口，该接口只能查询未完结的履约单以及完结的3天内订单
*/
type TmallCityretailWmflOrderLogisticsQueryRequest struct {
    model.Params
    // 订单号
    _mainOrderId   string
}

// 初始化TmallCityretailWmflOrderLogisticsQueryRequest对象
func NewTmallCityretailWmflOrderLogisticsQueryRequest() *TmallCityretailWmflOrderLogisticsQueryRequest{
    return &TmallCityretailWmflOrderLogisticsQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCityretailWmflOrderLogisticsQueryRequest) GetApiMethodName() string {
    return "tmall.cityretail.wmfl.order.logistics.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallCityretailWmflOrderLogisticsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 订单号
func (r *TmallCityretailWmflOrderLogisticsQueryRequest) SetMainOrderId(_mainOrderId string) error {
    r._mainOrderId = _mainOrderId
    r.Set("main_order_id", _mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TmallCityretailWmflOrderLogisticsQueryRequest) GetMainOrderId() string {
    return r._mainOrderId
}

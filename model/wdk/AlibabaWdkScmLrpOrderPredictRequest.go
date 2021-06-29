package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单量预测 API请求
alibaba.wdk.scm.lrp.order.predict

提供基于门店和配送站的订单量预测，可用于提前安排人力资源
*/
type AlibabaWdkScmLrpOrderPredictRequest struct {
    model.Params
    // 单量预测查询参数
    _paramOrderPredictQuery   *OrderPredictQuery
}

// 初始化AlibabaWdkScmLrpOrderPredictRequest对象
func NewAlibabaWdkScmLrpOrderPredictRequest() *AlibabaWdkScmLrpOrderPredictRequest{
    return &AlibabaWdkScmLrpOrderPredictRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkScmLrpOrderPredictRequest) GetApiMethodName() string {
    return "alibaba.wdk.scm.lrp.order.predict"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkScmLrpOrderPredictRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamOrderPredictQuery Setter
// 单量预测查询参数
func (r *AlibabaWdkScmLrpOrderPredictRequest) SetParamOrderPredictQuery(_paramOrderPredictQuery *OrderPredictQuery) error {
    r._paramOrderPredictQuery = _paramOrderPredictQuery
    r.Set("param_order_predict_query", _paramOrderPredictQuery)
    return nil
}

// ParamOrderPredictQuery Getter
func (r AlibabaWdkScmLrpOrderPredictRequest) GetParamOrderPredictQuery() *OrderPredictQuery {
    return r._paramOrderPredictQuery
}

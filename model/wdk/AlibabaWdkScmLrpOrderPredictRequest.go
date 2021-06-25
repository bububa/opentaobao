package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单量预测 APIRequest
alibaba.wdk.scm.lrp.order.predict

提供基于门店和配送站的订单量预测，可用于提前安排人力资源
*/
type AlibabaWdkScmLrpOrderPredictRequest struct {
    model.Params

    // 单量预测查询参数
    paramOrderPredictQuery   *OrderPredictQuery 

}

func NewAlibabaWdkScmLrpOrderPredictRequest() *AlibabaWdkScmLrpOrderPredictRequest{
    return &AlibabaWdkScmLrpOrderPredictRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkScmLrpOrderPredictRequest) GetApiMethodName() string {
    return "alibaba.wdk.scm.lrp.order.predict"
}

func (r AlibabaWdkScmLrpOrderPredictRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkScmLrpOrderPredictRequest) SetParamOrderPredictQuery(paramOrderPredictQuery *OrderPredictQuery) error {
    r.paramOrderPredictQuery = paramOrderPredictQuery
    r.Set("param_order_predict_query", paramOrderPredictQuery)
    return nil
}

func (r AlibabaWdkScmLrpOrderPredictRequest) GetParamOrderPredictQuery() *OrderPredictQuery {
    return r.paramOrderPredictQuery
}


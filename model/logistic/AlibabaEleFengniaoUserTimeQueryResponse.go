package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
蜂鸟询用户T APIResponse
alibaba.ele.fengniao.user.time.query

蜂鸟询用户T
*/
type AlibabaEleFengniaoUserTimeQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaEleFengniaoUserTimeQueryResponse `json:"alibaba_ele_fengniao_user_time_query_response,omitempty"`
}

type AlibabaEleFengniaoUserTimeQueryResponse struct {

    // 骑手预计时间
    CarrierPredictDeliveryTime   int64 `json:"carrier_predict_delivery_time,omitempty"`

    // 用户预计时间
    CustomerPredictDeliveryTime   int64 `json:"customer_predict_delivery_time,omitempty"`

}

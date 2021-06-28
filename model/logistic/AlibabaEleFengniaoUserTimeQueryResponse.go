package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
蜂鸟询用户T APIResponse
alibaba.ele.fengniao.user.time.query

蜂鸟询用户T
*/
type AlibabaEleFengniaoUserTimeQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_ele_fengniao_user_time_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 骑手预计时间
    
    CarrierPredictDeliveryTime   int64 `json:"carrier_predict_delivery_time,omitempty" xml:"
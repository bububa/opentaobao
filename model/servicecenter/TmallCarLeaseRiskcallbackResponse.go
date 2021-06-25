package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
整车租赁风控模型回调 APIResponse
tmall.car.lease.riskcallback

租赁公司回调风控结果
*/
type TmallCarLeaseRiskcallbackAPIResponse struct {
    model.CommonResponse
    Response *TmallCarLeaseRiskcallbackResponse `json:"tmall_car_lease_riskcallback_response,omitempty"`
}

type TmallCarLeaseRiskcallbackResponse struct {

    // 结果集合
    Result   *TmallCarLeaseRiskcallbackResult `json:"result,omitempty"`

}

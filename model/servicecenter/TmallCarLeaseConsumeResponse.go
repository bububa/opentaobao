package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
汽车租赁核销 APIResponse
tmall.car.lease.consume

租赁公司回传信息，核销
*/
type TmallCarLeaseConsumeAPIResponse struct {
    model.CommonResponse
    // Response *TmallCarLeaseConsumeResponse `json:"tmall_car_lease_consume_response,omitempty"` 
    TmallCarLeaseConsumeResponse
}

/* model for simplify = false
type TmallCarLeaseConsumeResponse struct {

    // 结果集合
    
    Result  *struct {
        TmallCarLeaseConsumeResult  *TmallCarLeaseConsumeResult `json:"tmall_car_lease_consume_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallCarLeaseConsumeResponse struct {

    // 结果集合
    Result   *TmallCarLeaseConsumeResult `json:"result,omitempty"`

}

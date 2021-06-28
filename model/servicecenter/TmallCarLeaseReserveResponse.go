package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
整车租车回传预约信息 APIResponse
tmall.car.lease.reserve

租赁公司回传预约到店信息
*/
type TmallCarLeaseReserveAPIResponse struct {
    model.CommonResponse
    // Response *TmallCarLeaseReserveResponse `json:"tmall_car_lease_reserve_response,omitempty"` 
    TmallCarLeaseReserveResponse
}

/* model for simplify = false
type TmallCarLeaseReserveResponse struct {

    // 返回结果
    
    Result  *struct {
        TmallCarLeaseReserveResult  *TmallCarLeaseReserveResult `json:"tmall_car_lease_reserve_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallCarLeaseReserveResponse struct {

    // 返回结果
    Result   *TmallCarLeaseReserveResult `json:"result,omitempty"`

}

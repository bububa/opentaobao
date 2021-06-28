package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
同步租赁方案 APIResponse
tmall.car.lease.synchronizeplans

租赁公司同步还款计划
*/
type TmallCarLeaseSynchronizeplansAPIResponse struct {
    model.CommonResponse
    // Response *TmallCarLeaseSynchronizeplansResponse `json:"tmall_car_lease_synchronizeplans_response,omitempty"` 
    TmallCarLeaseSynchronizeplansResponse
}

/* model for simplify = false
type TmallCarLeaseSynchronizeplansResponse struct {

    // 如果success为true,表示成功，如果success为false,需要获取msg_code,msg_info,具体的错误码文档提供
    
    Result  *struct {
        TmallCarLeaseSynchronizeplansResult  *TmallCarLeaseSynchronizeplansResult `json:"tmall_car_lease_synchronizeplans_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallCarLeaseSynchronizeplansResponse struct {

    // 如果success为true,表示成功，如果success为false,需要获取msg_code,msg_info,具体的错误码文档提供
    Result   *TmallCarLeaseSynchronizeplansResult `json:"result,omitempty"`

}

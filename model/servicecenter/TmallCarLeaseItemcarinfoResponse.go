package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
整车租赁商品四级车型信息 APIResponse
tmall.car.lease.itemcarinfo

整车租赁项目发布宝贝需要4级车型库，4级车型库信息需要回传
*/
type TmallCarLeaseItemcarinfoAPIResponse struct {
    model.CommonResponse
    // Response *TmallCarLeaseItemcarinfoResponse `json:"tmall_car_lease_itemcarinfo_response,omitempty"` 
    TmallCarLeaseItemcarinfoResponse
}

/* model for simplify = false
type TmallCarLeaseItemcarinfoResponse struct {

    // result
    
    Result  *struct {
        TmallCarLeaseItemcarinfoResult  *TmallCarLeaseItemcarinfoResult `json:"tmall_car_lease_itemcarinfo_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallCarLeaseItemcarinfoResponse struct {

    // result
    Result   *TmallCarLeaseItemcarinfoResult `json:"result,omitempty"`

}

package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询汽车租赁活动信息 APIResponse
tmall.car.lease.item.activity.get

查询汽车租赁活动信息
*/
type TmallCarLeaseItemActivityGetAPIResponse struct {
    model.CommonResponse
    Response *TmallCarLeaseItemActivityGetResponse `json:"tmall_car_lease_item_activity_get_response,omitempty"`
}

type TmallCarLeaseItemActivityGetResponse struct {

    // 结果对象
    Result   *TmallCarLeaseItemActivityGetResult `json:"result,omitempty"`

}

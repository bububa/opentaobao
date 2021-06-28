package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询汽车租赁活动信息 APIResponse
tmall.car.lease.item.activity.get

查询汽车租赁活动信息
*/
type TmallCarLeaseItemActivityGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_car_lease_item_activity_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果对象
    
    Result   *TmallCarLeaseItemActivityGetResult `json:"result,omitempty" xml:"
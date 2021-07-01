package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询汽车租赁活动信息 API返回值 
tmall.car.lease.item.activity.get

查询汽车租赁活动信息
*/
type TmallCarLeaseItemActivityGetAPIResponse struct {
    model.CommonResponse
    TmallCarLeaseItemActivityGetAPIResponseModel
}

// 查询汽车租赁活动信息 成功返回结果
type TmallCarLeaseItemActivityGetAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_car_lease_item_activity_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果对象
    Result   *TmallCarLeaseItemActivityGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

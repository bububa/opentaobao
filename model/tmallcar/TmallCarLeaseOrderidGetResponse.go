package tmallcar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车查询订单id API返回值 
tmall.car.lease.orderid.get

天猫开新车查询订单id
*/
type TmallCarLeaseOrderidGetAPIResponse struct {
    model.CommonResponse
    TmallCarLeaseOrderidGetResponse
}

// 天猫开新车查询订单id 成功返回结果
type TmallCarLeaseOrderidGetResponse struct {
    XMLName xml.Name `xml:"tmall_car_lease_orderid_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}

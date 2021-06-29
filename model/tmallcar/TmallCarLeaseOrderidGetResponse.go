package tmallcar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车查询订单id APIResponse
tmall.car.lease.orderid.get

天猫开新车查询订单id
*/
type TmallCarLeaseOrderidGetAPIResponse struct {
    model.CommonResponse
    TmallCarLeaseOrderidGetResponse
}

type TmallCarLeaseOrderidGetResponse struct {
    XMLName xml.Name `xml:"tmall_car_lease_orderid_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ResultVo `json:"result,omitempty" xml:"result,omitempty"`

    
}

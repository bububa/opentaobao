package tmallcar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车租后代客户还款 APIResponse
tmall.car.lease.payforcustomer

天猫开新车租后代客户还款
*/
type TmallCarLeasePayforcustomerAPIResponse struct {
    model.CommonResponse
    TmallCarLeasePayforcustomerResponse
}

type TmallCarLeasePayforcustomerResponse struct {
    XMLName xml.Name `xml:"tmall_car_lease_payforcustomer_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ResultVo `json:"result,omitempty" xml:"result,omitempty"`

    
}

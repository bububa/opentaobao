package tmallcar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车租后查询还款计划 APIResponse
tmall.car.lease.queryloanplans

天猫开新车租后查询还款计划
*/
type TmallCarLeaseQueryloanplansAPIResponse struct {
    model.CommonResponse
    TmallCarLeaseQueryloanplansResponse
}

type TmallCarLeaseQueryloanplansResponse struct {
    XMLName xml.Name `xml:"tmall_car_lease_queryloanplans_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *ResultVo `json:"result,omitempty" xml:"result,omitempty"`

    
}

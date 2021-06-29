package tmallcar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车租后异常流线下处理状态通知接口 APIResponse
tmall.car.lease.exceptionflowsynchronize

天猫开新车租后异常流线下处理状态通知接口
*/
type TmallCarLeaseExceptionflowsynchronizeAPIResponse struct {
    model.CommonResponse
    TmallCarLeaseExceptionflowsynchronizeResponse
}

type TmallCarLeaseExceptionflowsynchronizeResponse struct {
    XMLName xml.Name `xml:"tmall_car_lease_exceptionflowsynchronize_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ResultVo `json:"result,omitempty" xml:"result,omitempty"`

    
}

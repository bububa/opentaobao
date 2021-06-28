package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车租赁核销 APIResponse
tmall.car.lease.consume

租赁公司回传信息，核销
*/
type TmallCarLeaseConsumeAPIResponse struct {
    model.CommonResponse
    TmallCarLeaseConsumeResponse
}

type TmallCarLeaseConsumeResponse struct {
    XMLName xml.Name `xml:"tmall_car_lease_consume_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果集合
    
    Result   *TmallCarLeaseConsumeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

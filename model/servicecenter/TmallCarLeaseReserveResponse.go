package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
整车租车回传预约信息 APIResponse
tmall.car.lease.reserve

租赁公司回传预约到店信息
*/
type TmallCarLeaseReserveAPIResponse struct {
    model.CommonResponse
    TmallCarLeaseReserveResponse
}

type TmallCarLeaseReserveResponse struct {
    XMLName xml.Name `xml:"tmall_car_lease_reserve_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TmallCarLeaseReserveResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

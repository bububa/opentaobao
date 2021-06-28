package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
整车租赁风控模型回调 APIResponse
tmall.car.lease.riskcallback

租赁公司回调风控结果
*/
type TmallCarLeaseRiskcallbackAPIResponse struct {
    model.CommonResponse
    TmallCarLeaseRiskcallbackResponse
}

type TmallCarLeaseRiskcallbackResponse struct {
    XMLName xml.Name `xml:"tmall_car_lease_riskcallback_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果集合
    
    Result   *TmallCarLeaseRiskcallbackResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

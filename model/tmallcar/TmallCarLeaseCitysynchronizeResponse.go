package tmallcar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车租后分期城市信息同步 APIResponse
tmall.car.lease.citysynchronize

天猫开新车租后分期城市信息同步
*/
type TmallCarLeaseCitysynchronizeAPIResponse struct {
    model.CommonResponse
    TmallCarLeaseCitysynchronizeResponse
}

type TmallCarLeaseCitysynchronizeResponse struct {
    XMLName xml.Name `xml:"tmall_car_lease_citysynchronize_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ResultVo `json:"result,omitempty" xml:"result,omitempty"`

    
}

package tmallcar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车租后合同下载 APIResponse
tmall.car.lease.contractdownload

天猫开新车租后合同下载
*/
type TmallCarLeaseContractdownloadAPIResponse struct {
    model.CommonResponse
    TmallCarLeaseContractdownloadResponse
}

type TmallCarLeaseContractdownloadResponse struct {
    XMLName xml.Name `xml:"tmall_car_lease_contractdownload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *ResultVo `json:"result,omitempty" xml:"result,omitempty"`

    
}

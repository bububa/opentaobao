package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增成本中心 APIResponse
alitrip.btrip.open.cost.center.new

新增成本中心
*/
type AlitripBtripOpenCostCenterNewAPIResponse struct {
    model.CommonResponse
    AlitripBtripOpenCostCenterNewResponse
}

type AlitripBtripOpenCostCenterNewResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_open_cost_center_new_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 成本中心对象
    
    Module   *OpenCostCenterSaveRs `json:"module,omitempty" xml:"module,omitempty"`

    
    // 结果码
    
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 结果描述
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
}

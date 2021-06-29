package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
增加成本中心人员信息 APIResponse
alitrip.btrip.open.cost.center.entity.add

增加成本中心人员信息
*/
type AlitripBtripOpenCostCenterEntityAddAPIResponse struct {
    model.CommonResponse
    AlitripBtripOpenCostCenterEntityAddResponse
}

type AlitripBtripOpenCostCenterEntityAddResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_open_cost_center_entity_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果对象
    
    Result   *OpenCostCenterAddEntityRs `json:"result,omitempty" xml:"result,omitempty"`

    
    // 结果码
    
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 结果描述
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
}

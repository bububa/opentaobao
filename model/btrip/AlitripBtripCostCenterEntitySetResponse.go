package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设置外部成本中心人员信息 APIResponse
alitrip.btrip.cost.center.entity.set

设置外部成本中心人员信息
*/
type AlitripBtripCostCenterEntitySetAPIResponse struct {
    model.CommonResponse
    AlitripBtripCostCenterEntitySetResponse
}

type AlitripBtripCostCenterEntitySetResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_cost_center_entity_set_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象
    
    Result   *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

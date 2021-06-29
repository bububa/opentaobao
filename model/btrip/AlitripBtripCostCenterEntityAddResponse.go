package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
增加外部成本中心人员信息 APIResponse
alitrip.btrip.cost.center.entity.add

增加外部成本中心人员信息
*/
type AlitripBtripCostCenterEntityAddAPIResponse struct {
    model.CommonResponse
    AlitripBtripCostCenterEntityAddResponse
}

type AlitripBtripCostCenterEntityAddResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_cost_center_entity_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象
    
    Result   *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

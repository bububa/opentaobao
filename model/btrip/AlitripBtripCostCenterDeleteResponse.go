package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除外部成本中心 APIResponse
alitrip.btrip.cost.center.delete

删除外部成本中心
*/
type AlitripBtripCostCenterDeleteAPIResponse struct {
    model.CommonResponse
    AlitripBtripCostCenterDeleteResponse
}

type AlitripBtripCostCenterDeleteResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_cost_center_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象
    
    Result   *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

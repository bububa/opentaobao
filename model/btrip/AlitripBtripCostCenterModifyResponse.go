package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改外部成本中心 APIResponse
alitrip.btrip.cost.center.modify

修改外部成本中心，设置成员，设置支付宝账号，设置名称，编号等
*/
type AlitripBtripCostCenterModifyAPIResponse struct {
    model.CommonResponse
    AlitripBtripCostCenterModifyResponse
}

type AlitripBtripCostCenterModifyResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_cost_center_modify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象
    
    Result   *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

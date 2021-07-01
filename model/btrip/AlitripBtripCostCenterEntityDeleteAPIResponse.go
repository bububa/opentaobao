package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除外部成本中心人员信息 API返回值 
alitrip.btrip.cost.center.entity.delete

删除外部成本中心人员信息
*/
type AlitripBtripCostCenterEntityDeleteAPIResponse struct {
    model.CommonResponse
    AlitripBtripCostCenterEntityDeleteAPIResponseModel
}

// 删除外部成本中心人员信息 成功返回结果
type AlitripBtripCostCenterEntityDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_btrip_cost_center_entity_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回对象
    Result   *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

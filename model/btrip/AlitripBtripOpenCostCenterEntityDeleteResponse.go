package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除成本中心人员信息 API返回值 
alitrip.btrip.open.cost.center.entity.delete

删除成本中心人员信息
*/
type AlitripBtripOpenCostCenterEntityDeleteAPIResponse struct {
    model.CommonResponse
    AlitripBtripOpenCostCenterEntityDeleteResponse
}

// 删除成本中心人员信息 成功返回结果
type AlitripBtripOpenCostCenterEntityDeleteResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_open_cost_center_entity_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果对象
    Result   *OpenCostCenterDeleteEntityRs `json:"result,omitempty" xml:"result,omitempty"`
    // 结果码
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 结果信息
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}

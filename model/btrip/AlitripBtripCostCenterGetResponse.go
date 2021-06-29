package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户费用归属 API返回值 
alitrip.btrip.cost.center.get

获取差旅申请用户的费用归属列表
*/
type AlitripBtripCostCenterGetAPIResponse struct {
    model.CommonResponse
    AlitripBtripCostCenterGetResponse
}

// 获取用户费用归属 成功返回结果
type AlitripBtripCostCenterGetResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_cost_center_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

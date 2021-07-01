package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询外部成本中心 API返回值 
alitrip.btip.cost.center.query

查询外部成本中心
*/
type AlitripBtipCostCenterQueryAPIResponse struct {
    model.CommonResponse
    AlitripBtipCostCenterQueryAPIResponseModel
}

// 查询外部成本中心 成功返回结果
type AlitripBtipCostCenterQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_btip_cost_center_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回对象
    Result   *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

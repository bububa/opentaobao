package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripOpenCostCenterQueryAPIResponse
查询成本中心 API返回值
alitrip.btrip.open.cost.center.query

查询成本中心 */
type AlitripBtripOpenCostCenterQueryAPIResponse struct {
	model.CommonResponse
	AlitripBtripOpenCostCenterQueryAPIResponseModel
}

// AlitripBtripOpenCostCenterQueryAPIResponseModel is 查询成本中心 成功返回结果
type AlitripBtripOpenCostCenterQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_open_cost_center_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成本中心列表
	CostCenterList []OpenCostCenterQueryRs `json:"cost_center_list,omitempty" xml:"cost_center_list>open_cost_center_query_rs,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}

package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterDeleteAPIResponse 删除成本中心 API返回值
// alitrip.btrip.open.cost.center.delete
//
// 删除成本中心
type AlitripBtripOpenCostCenterDeleteAPIResponse struct {
	model.CommonResponse
	AlitripBtripOpenCostCenterDeleteAPIResponseModel
}

// AlitripBtripOpenCostCenterDeleteAPIResponseModel is 删除成本中心 成功返回结果
type AlitripBtripOpenCostCenterDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_open_cost_center_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

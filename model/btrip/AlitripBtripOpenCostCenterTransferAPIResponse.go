package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterTransferAPIResponse 商旅成本中心转换为外部成本中心 API返回值
// alitrip.btrip.open.cost.center.transfer
//
// 商旅成本中心转换为外部成本中心
type AlitripBtripOpenCostCenterTransferAPIResponse struct {
	model.CommonResponse
	AlitripBtripOpenCostCenterTransferAPIResponseModel
}

// AlitripBtripOpenCostCenterTransferAPIResponseModel is 商旅成本中心转换为外部成本中心 成功返回结果
type AlitripBtripOpenCostCenterTransferAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_open_cost_center_transfer_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}

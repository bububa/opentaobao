package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCostCenterDeleteAPIResponse
删除外部成本中心 API返回值
alitrip.btrip.cost.center.delete

删除外部成本中心 */
type AlitripBtripCostCenterDeleteAPIResponse struct {
	model.CommonResponse
	AlitripBtripCostCenterDeleteAPIResponseModel
}

// AlitripBtripCostCenterDeleteAPIResponseModel is 删除外部成本中心 成功返回结果
type AlitripBtripCostCenterDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_cost_center_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

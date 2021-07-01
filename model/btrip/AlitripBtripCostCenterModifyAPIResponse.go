package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCostCenterModifyAPIResponse
修改外部成本中心 API返回值
alitrip.btrip.cost.center.modify

修改外部成本中心，设置成员，设置支付宝账号，设置名称，编号等 */
type AlitripBtripCostCenterModifyAPIResponse struct {
	model.CommonResponse
	AlitripBtripCostCenterModifyAPIResponseModel
}

// AlitripBtripCostCenterModifyAPIResponseModel is 修改外部成本中心 成功返回结果
type AlitripBtripCostCenterModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_cost_center_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

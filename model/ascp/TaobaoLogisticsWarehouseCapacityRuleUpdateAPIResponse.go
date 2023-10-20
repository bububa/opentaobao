package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswarehousecapacityruleupdateAPIResponse 仓产能创建/更新 API返回值
// taobao.logistics.warehouse.capacity.rule.update
//
// 仓产能创建/更新
type TaobaologisticswarehousecapacityruleupdateAPIResponse struct {
	model.CommonResponse
	TaobaologisticswarehousecapacityruleupdateAPIResponseModel
}

// TaobaologisticswarehousecapacityruleupdateAPIResponseModel is 仓产能创建/更新 成功返回结果
type TaobaologisticswarehousecapacityruleupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_warehouse_capacity_rule_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 仓产能创建/更新出参
	CapacityRuleUpdateResponse *CapacityRuleUpdateResponse `json:"capacity_rule_update_response,omitempty" xml:"capacity_rule_update_response,omitempty"`
}

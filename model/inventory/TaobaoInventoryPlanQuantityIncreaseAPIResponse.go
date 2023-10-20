package inventory

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoinventoryplanquantityincreaseAPIResponse 计划库存的增量编辑 API返回值
// taobao.inventory.plan.quantity.increase
//
// 计划库存的增量编辑
type TaobaoinventoryplanquantityincreaseAPIResponse struct {
	model.CommonResponse
	TaobaoinventoryplanquantityincreaseAPIResponseModel
}

// TaobaoinventoryplanquantityincreaseAPIResponseModel is 计划库存的增量编辑 成功返回结果
type TaobaoinventoryplanquantityincreaseAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_plan_quantity_increase_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 批量返回结果
	Result *BatchResult `json:"result,omitempty" xml:"result,omitempty"`
}

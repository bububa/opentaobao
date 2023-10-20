package inventory

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoinventoryplaneditAPIResponse 设置计划库存 API返回值
// taobao.inventory.plan.edit
//
// 初始化计划库存，或者编辑已经存在的计划库存
type TaobaoinventoryplaneditAPIResponse struct {
	model.CommonResponse
	TaobaoinventoryplaneditAPIResponseModel
}

// TaobaoinventoryplaneditAPIResponseModel is 设置计划库存 成功返回结果
type TaobaoinventoryplaneditAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_plan_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 批量返回结果
	Result *BatchResult `json:"result,omitempty" xml:"result,omitempty"`
}

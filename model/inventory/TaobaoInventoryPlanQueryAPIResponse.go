package inventory

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoinventoryplanqueryAPIResponse 计划库存查询 API返回值
// taobao.inventory.plan.query
//
// 计划库存查询
type TaobaoinventoryplanqueryAPIResponse struct {
	model.CommonResponse
	TaobaoinventoryplanqueryAPIResponseModel
}

// TaobaoinventoryplanqueryAPIResponseModel is 计划库存查询 成功返回结果
type TaobaoinventoryplanqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_plan_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoinventoryplanqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

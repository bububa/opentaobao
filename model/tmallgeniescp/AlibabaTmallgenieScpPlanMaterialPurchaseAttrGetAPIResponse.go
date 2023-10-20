package tmallgeniescp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIResponse 物料的采购属性查询 API返回值
// alibaba.tmallgenie.scp.plan.material.purchase.attr.get
//
// 物料的采购属性查询
type AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIResponse struct {
	model.CommonResponse
	AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIResponseModel
}

// AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIResponseModel is 物料的采购属性查询 成功返回结果
type AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_material_purchase_attr_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

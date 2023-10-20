package tmallgeniescp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse 物料的采购属性查询 API返回值
// alibaba.tmallgenie.scp.plan.material.purchase.attr.get
//
// 物料的采购属性查询
type AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponseModel).Reset()
}

// AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponseModel is 物料的采购属性查询 成功返回结果
type AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_material_purchase_attr_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse)
	},
}

// GetAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse 从 sync.Pool 获取 AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse
func GetAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse() *AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse {
	return poolAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse.Get().(*AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse)
}

// ReleaseAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse 将 AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse(v *AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse.Put(v)
}

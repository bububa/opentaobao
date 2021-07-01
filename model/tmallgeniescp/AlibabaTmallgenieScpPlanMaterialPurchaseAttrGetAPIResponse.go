package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物料的采购属性查询 API返回值 
alibaba.tmallgenie.scp.plan.material.purchase.attr.get

物料的采购属性查询
*/
type AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponseModel
}

// 物料的采购属性查询 成功返回结果
type AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_material_purchase_attr_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回对象封装
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物料的采购属性查询 APIResponse
alibaba.tmallgenie.scp.plan.material.purchase.attr.get

物料的采购属性查询
*/
type AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetResponse
}

type AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetResponse struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_material_purchase_attr_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象封装
    
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

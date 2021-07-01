package admarket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
广告平台创意审核 API返回值 
yunos.admarket.material.audit

用于厂商上报广告平台审核结果
*/
type YunosAdmarketMaterialAuditAPIResponse struct {
    model.CommonResponse
    YunosAdmarketMaterialAuditAPIResponseModel
}

// 广告平台创意审核 成功返回结果
type YunosAdmarketMaterialAuditAPIResponseModel struct {
    XMLName xml.Name `xml:"yunos_admarket_material_audit_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *YunosAdmarketMaterialAuditResult `json:"result,omitempty" xml:"result,omitempty"`
}

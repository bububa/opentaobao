package admarket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
广告平台创意审核 APIRequest
yunos.admarket.material.audit

用于厂商上报广告平台审核结果
*/
type YunosAdmarketMaterialAuditRequest struct {
    model.Params

    // 创意审核结果
    sspMaterialAuditResult   *SspMaterialAuditResult 

}

func NewYunosAdmarketMaterialAuditRequest() *YunosAdmarketMaterialAuditRequest{
    return &YunosAdmarketMaterialAuditRequest{
        Params: model.NewParams(),
    }
}

func (r YunosAdmarketMaterialAuditRequest) GetApiMethodName() string {
    return "yunos.admarket.material.audit"
}

func (r YunosAdmarketMaterialAuditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosAdmarketMaterialAuditRequest) SetSspMaterialAuditResult(sspMaterialAuditResult *SspMaterialAuditResult) error {
    r.sspMaterialAuditResult = sspMaterialAuditResult
    r.Set("ssp_material_audit_result", sspMaterialAuditResult)
    return nil
}

func (r YunosAdmarketMaterialAuditRequest) GetSspMaterialAuditResult() *SspMaterialAuditResult {
    return r.sspMaterialAuditResult
}


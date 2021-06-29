package admarket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/admarket"
)

/* 
广告平台创意审核 
yunos.admarket.material.audit

用于厂商上报广告平台审核结果
*/
func YunosAdmarketMaterialAudit(clt *core.SDKClient, req *admarket.YunosAdmarketMaterialAuditRequest, session string) (*admarket.YunosAdmarketMaterialAuditAPIResponse, error) {
    var resp admarket.YunosAdmarketMaterialAuditAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

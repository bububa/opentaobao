package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物料的采购属性查询 APIRequest
alibaba.tmallgenie.scp.plan.material.purchase.attr.get

物料的采购属性查询
*/
type AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest struct {
    model.Params

    // 扩展字段
    requestExtendJson   string 

}

func NewAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest() *AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest{
    return &AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.material.purchase.attr.get"
}

func (r AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest) SetRequestExtendJson(requestExtendJson string) error {
    r.requestExtendJson = requestExtendJson
    r.Set("request_extend_json", requestExtendJson)
    return nil
}

func (r AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest) GetRequestExtendJson() string {
    return r.requestExtendJson
}


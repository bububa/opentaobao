package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物料的采购属性查询 API请求
alibaba.tmallgenie.scp.plan.material.purchase.attr.get

物料的采购属性查询
*/
type AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest struct {
    model.Params
    // 扩展字段
    _requestExtendJson   string
}

// 初始化AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest对象
func NewAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest() *AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest{
    return &AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.material.purchase.attr.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestExtendJson Setter
// 扩展字段
func (r *AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest) SetRequestExtendJson(_requestExtendJson string) error {
    r._requestExtendJson = _requestExtendJson
    r.Set("request_extend_json", _requestExtendJson)
    return nil
}

// RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest) GetRequestExtendJson() string {
    return r._requestExtendJson
}

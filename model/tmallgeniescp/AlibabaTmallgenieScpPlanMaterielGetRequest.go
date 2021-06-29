package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
1-IBP同步物料接口 API请求
alibaba.tmallgenie.scp.plan.materiel.get

IBP同步物料接口
*/
type AlibabaTmallgenieScpPlanMaterielGetRequest struct {
    model.Params
    // 扩展字段
    requestExtendJson   string
}

// 初始化AlibabaTmallgenieScpPlanMaterielGetRequest对象
func NewAlibabaTmallgenieScpPlanMaterielGetRequest() *AlibabaTmallgenieScpPlanMaterielGetRequest{
    return &AlibabaTmallgenieScpPlanMaterielGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanMaterielGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.materiel.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanMaterielGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestExtendJson Setter
// 扩展字段
func (r *AlibabaTmallgenieScpPlanMaterielGetRequest) SetRequestExtendJson(requestExtendJson string) error {
    r.requestExtendJson = requestExtendJson
    r.Set("request_extend_json", requestExtendJson)
    return nil
}

// RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanMaterielGetRequest) GetRequestExtendJson() string {
    return r.requestExtendJson
}

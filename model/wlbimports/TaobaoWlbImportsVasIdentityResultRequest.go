package wlbimports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
集货鉴定结果 APIRequest
taobao.wlb.imports.vas.identity.result

集货鉴定结果查询
*/
type TaobaoWlbImportsVasIdentityResultRequest struct {
    model.Params

    // 物流订单编号
    lgOrderCode   string 

}

func NewTaobaoWlbImportsVasIdentityResultRequest() *TaobaoWlbImportsVasIdentityResultRequest{
    return &TaobaoWlbImportsVasIdentityResultRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbImportsVasIdentityResultRequest) GetApiMethodName() string {
    return "taobao.wlb.imports.vas.identity.result"
}

func (r TaobaoWlbImportsVasIdentityResultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbImportsVasIdentityResultRequest) SetLgOrderCode(lgOrderCode string) error {
    r.lgOrderCode = lgOrderCode
    r.Set("lg_order_code", lgOrderCode)
    return nil
}

func (r TaobaoWlbImportsVasIdentityResultRequest) GetLgOrderCode() string {
    return r.lgOrderCode
}


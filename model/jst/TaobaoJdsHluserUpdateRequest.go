package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单全链路用户信息修改 APIRequest
taobao.jds.hluser.update

订单全链路用户信息修改，比如是否开放买家端展示
*/
type TaobaoJdsHluserUpdateRequest struct {
    model.Params

    // 回流信息是否开通买家端展示,可选值open,close
    openForBuyer   string 

    // 为空,则默认是X_TO_SYSTEM,X_WAIT_ALLOCATION,X_OUT_WAREHOUSE 可选值是X_DOWNLOADED,X_TO_SYSTEM,X_SERVICE_AUDITED,X_FINANCE_AUDITED,X_ALLOCATION_NOTIFIED,X_WAIT_ALLOCATION,X_SORT_PRINTED,X_SEND_PRINTED,X_LOGISTICS_PRINTED,X_SORTED,X_EXAMINED,X_PACKAGED,X_WEIGHED,X_OUT_WAREHOUS
    openNodes   string 

}

func NewTaobaoJdsHluserUpdateRequest() *TaobaoJdsHluserUpdateRequest{
    return &TaobaoJdsHluserUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJdsHluserUpdateRequest) GetApiMethodName() string {
    return "taobao.jds.hluser.update"
}

func (r TaobaoJdsHluserUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJdsHluserUpdateRequest) SetOpenForBuyer(openForBuyer string) error {
    r.openForBuyer = openForBuyer
    r.Set("open_for_buyer", openForBuyer)
    return nil
}

func (r TaobaoJdsHluserUpdateRequest) GetOpenForBuyer() string {
    return r.openForBuyer
}

func (r *TaobaoJdsHluserUpdateRequest) SetOpenNodes(openNodes string) error {
    r.openNodes = openNodes
    r.Set("open_nodes", openNodes)
    return nil
}

func (r TaobaoJdsHluserUpdateRequest) GetOpenNodes() string {
    return r.openNodes
}


package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单全链路用户信息修改 API请求
taobao.jds.hluser.update

订单全链路用户信息修改，比如是否开放买家端展示
*/
type TaobaoJdsHluserUpdateRequest struct {
    model.Params
    // 回流信息是否开通买家端展示,可选值open,close
    _openForBuyer   string
    // 为空,则默认是X_TO_SYSTEM,X_WAIT_ALLOCATION,X_OUT_WAREHOUSE 可选值是X_DOWNLOADED,X_TO_SYSTEM,X_SERVICE_AUDITED,X_FINANCE_AUDITED,X_ALLOCATION_NOTIFIED,X_WAIT_ALLOCATION,X_SORT_PRINTED,X_SEND_PRINTED,X_LOGISTICS_PRINTED,X_SORTED,X_EXAMINED,X_PACKAGED,X_WEIGHED,X_OUT_WAREHOUS
    _openNodes   string
}

// 初始化TaobaoJdsHluserUpdateRequest对象
func NewTaobaoJdsHluserUpdateRequest() *TaobaoJdsHluserUpdateRequest{
    return &TaobaoJdsHluserUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJdsHluserUpdateRequest) GetApiMethodName() string {
    return "taobao.jds.hluser.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJdsHluserUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenForBuyer Setter
// 回流信息是否开通买家端展示,可选值open,close
func (r *TaobaoJdsHluserUpdateRequest) SetOpenForBuyer(_openForBuyer string) error {
    r._openForBuyer = _openForBuyer
    r.Set("open_for_buyer", _openForBuyer)
    return nil
}

// OpenForBuyer Getter
func (r TaobaoJdsHluserUpdateRequest) GetOpenForBuyer() string {
    return r._openForBuyer
}
// OpenNodes Setter
// 为空,则默认是X_TO_SYSTEM,X_WAIT_ALLOCATION,X_OUT_WAREHOUSE 可选值是X_DOWNLOADED,X_TO_SYSTEM,X_SERVICE_AUDITED,X_FINANCE_AUDITED,X_ALLOCATION_NOTIFIED,X_WAIT_ALLOCATION,X_SORT_PRINTED,X_SEND_PRINTED,X_LOGISTICS_PRINTED,X_SORTED,X_EXAMINED,X_PACKAGED,X_WEIGHED,X_OUT_WAREHOUS
func (r *TaobaoJdsHluserUpdateRequest) SetOpenNodes(_openNodes string) error {
    r._openNodes = _openNodes
    r.Set("open_nodes", _openNodes)
    return nil
}

// OpenNodes Getter
func (r TaobaoJdsHluserUpdateRequest) GetOpenNodes() string {
    return r._openNodes
}

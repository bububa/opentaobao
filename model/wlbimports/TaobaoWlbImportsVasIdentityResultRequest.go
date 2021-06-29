package wlbimports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
集货鉴定结果 API请求
taobao.wlb.imports.vas.identity.result

集货鉴定结果查询
*/
type TaobaoWlbImportsVasIdentityResultRequest struct {
    model.Params
    // 物流订单编号
    _lgOrderCode   string
}

// 初始化TaobaoWlbImportsVasIdentityResultRequest对象
func NewTaobaoWlbImportsVasIdentityResultRequest() *TaobaoWlbImportsVasIdentityResultRequest{
    return &TaobaoWlbImportsVasIdentityResultRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportsVasIdentityResultRequest) GetApiMethodName() string {
    return "taobao.wlb.imports.vas.identity.result"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportsVasIdentityResultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LgOrderCode Setter
// 物流订单编号
func (r *TaobaoWlbImportsVasIdentityResultRequest) SetLgOrderCode(_lgOrderCode string) error {
    r._lgOrderCode = _lgOrderCode
    r.Set("lg_order_code", _lgOrderCode)
    return nil
}

// LgOrderCode Getter
func (r TaobaoWlbImportsVasIdentityResultRequest) GetLgOrderCode() string {
    return r._lgOrderCode
}

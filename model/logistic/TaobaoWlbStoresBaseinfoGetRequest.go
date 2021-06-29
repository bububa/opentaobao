package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家按照仓的类型获取仓库接口 API请求
taobao.wlb.stores.baseinfo.get

通过USERID和仓库类型，获取商家自有仓库或菜鸟仓库或全部仓库
*/
type TaobaoWlbStoresBaseinfoGetRequest struct {
    model.Params
    // 0.商家仓库.1.菜鸟仓库.2全部被划分的仓库
    _type   int64
}

// 初始化TaobaoWlbStoresBaseinfoGetRequest对象
func NewTaobaoWlbStoresBaseinfoGetRequest() *TaobaoWlbStoresBaseinfoGetRequest{
    return &TaobaoWlbStoresBaseinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbStoresBaseinfoGetRequest) GetApiMethodName() string {
    return "taobao.wlb.stores.baseinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbStoresBaseinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 0.商家仓库.1.菜鸟仓库.2全部被划分的仓库
func (r *TaobaoWlbStoresBaseinfoGetRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoWlbStoresBaseinfoGetRequest) GetType() int64 {
    return r._type
}

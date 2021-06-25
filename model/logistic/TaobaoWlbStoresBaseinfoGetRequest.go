package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家按照仓的类型获取仓库接口 APIRequest
taobao.wlb.stores.baseinfo.get

通过USERID和仓库类型，获取商家自有仓库或菜鸟仓库或全部仓库
*/
type TaobaoWlbStoresBaseinfoGetRequest struct {
    model.Params

    // 0.商家仓库.1.菜鸟仓库.2全部被划分的仓库
    type   int64 

}

func NewTaobaoWlbStoresBaseinfoGetRequest() *TaobaoWlbStoresBaseinfoGetRequest{
    return &TaobaoWlbStoresBaseinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbStoresBaseinfoGetRequest) GetApiMethodName() string {
    return "taobao.wlb.stores.baseinfo.get"
}

func (r TaobaoWlbStoresBaseinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbStoresBaseinfoGetRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoWlbStoresBaseinfoGetRequest) GetType() int64 {
    return r.type
}


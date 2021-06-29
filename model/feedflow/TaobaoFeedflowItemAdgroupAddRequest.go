package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流增加单元 APIRequest
taobao.feedflow.item.adgroup.add

信息流增加单元
*/
type TaobaoFeedflowItemAdgroupAddRequest struct {
    model.Params

    // 单元信息
    adgroup   *AdgroupDto 

}

func NewTaobaoFeedflowItemAdgroupAddRequest() *TaobaoFeedflowItemAdgroupAddRequest{
    return &TaobaoFeedflowItemAdgroupAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemAdgroupAddRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.add"
}

func (r TaobaoFeedflowItemAdgroupAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemAdgroupAddRequest) SetAdgroup(adgroup *AdgroupDto) error {
    r.adgroup = adgroup
    r.Set("adgroup", adgroup)
    return nil
}

func (r TaobaoFeedflowItemAdgroupAddRequest) GetAdgroup() *AdgroupDto {
    return r.adgroup
}


package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流单元修改 APIRequest
taobao.feedflow.item.adgroup.modify

信息流单元修改
*/
type TaobaoFeedflowItemAdgroupModifyRequest struct {
    model.Params

    // 单元信息
    adgroup   *AdgroupDto 

}

func NewTaobaoFeedflowItemAdgroupModifyRequest() *TaobaoFeedflowItemAdgroupModifyRequest{
    return &TaobaoFeedflowItemAdgroupModifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemAdgroupModifyRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.modify"
}

func (r TaobaoFeedflowItemAdgroupModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemAdgroupModifyRequest) SetAdgroup(adgroup *AdgroupDto) error {
    r.adgroup = adgroup
    r.Set("adgroup", adgroup)
    return nil
}

func (r TaobaoFeedflowItemAdgroupModifyRequest) GetAdgroup() *AdgroupDto {
    return r.adgroup
}


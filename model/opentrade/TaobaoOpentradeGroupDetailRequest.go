package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购场景查询团详情 APIRequest
taobao.opentrade.group.detail

组团购场景下，查询团详情
*/
type TaobaoOpentradeGroupDetailRequest struct {
    model.Params

    // 团id
    groupId   int64 

}

func NewTaobaoOpentradeGroupDetailRequest() *TaobaoOpentradeGroupDetailRequest{
    return &TaobaoOpentradeGroupDetailRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpentradeGroupDetailRequest) GetApiMethodName() string {
    return "taobao.opentrade.group.detail"
}

func (r TaobaoOpentradeGroupDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpentradeGroupDetailRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r TaobaoOpentradeGroupDetailRequest) GetGroupId() int64 {
    return r.groupId
}


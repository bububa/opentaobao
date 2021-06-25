package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除活动范围 APIRequest
taobao.ump.range.delete

去指先前指定在某项活动中，某些类型的物品参加或者不参加活动的设置
*/
type TaobaoUmpRangeDeleteRequest struct {
    model.Params

    // 活动id
    actId   int64 

    // 范围的类型，比如是全店，商品，类目见：MarketingConstants.PARTICIPATE_TYPE_*
    type   int64 

    // id列表，当范围类型为商品时，该id为商品id；当范围类型为类目时，该id为类目id
    ids   string 

}

func NewTaobaoUmpRangeDeleteRequest() *TaobaoUmpRangeDeleteRequest{
    return &TaobaoUmpRangeDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpRangeDeleteRequest) GetApiMethodName() string {
    return "taobao.ump.range.delete"
}

func (r TaobaoUmpRangeDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpRangeDeleteRequest) SetActId(actId int64) error {
    r.actId = actId
    r.Set("act_id", actId)
    return nil
}

func (r TaobaoUmpRangeDeleteRequest) GetActId() int64 {
    return r.actId
}

func (r *TaobaoUmpRangeDeleteRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoUmpRangeDeleteRequest) GetType() int64 {
    return r.type
}

func (r *TaobaoUmpRangeDeleteRequest) SetIds(ids string) error {
    r.ids = ids
    r.Set("ids", ids)
    return nil
}

func (r TaobaoUmpRangeDeleteRequest) GetIds() string {
    return r.ids
}


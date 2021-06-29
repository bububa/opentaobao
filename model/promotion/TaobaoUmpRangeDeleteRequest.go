package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除活动范围 API请求
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

// 初始化TaobaoUmpRangeDeleteRequest对象
func NewTaobaoUmpRangeDeleteRequest() *TaobaoUmpRangeDeleteRequest{
    return &TaobaoUmpRangeDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpRangeDeleteRequest) GetApiMethodName() string {
    return "taobao.ump.range.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpRangeDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActId Setter
// 活动id
func (r *TaobaoUmpRangeDeleteRequest) SetActId(actId int64) error {
    r.actId = actId
    r.Set("act_id", actId)
    return nil
}

// ActId Getter
func (r TaobaoUmpRangeDeleteRequest) GetActId() int64 {
    return r.actId
}
// Type Setter
// 范围的类型，比如是全店，商品，类目见：MarketingConstants.PARTICIPATE_TYPE_*
func (r *TaobaoUmpRangeDeleteRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoUmpRangeDeleteRequest) GetType() int64 {
    return r.type
}
// Ids Setter
// id列表，当范围类型为商品时，该id为商品id；当范围类型为类目时，该id为类目id
func (r *TaobaoUmpRangeDeleteRequest) SetIds(ids string) error {
    r.ids = ids
    r.Set("ids", ids)
    return nil
}

// Ids Getter
func (r TaobaoUmpRangeDeleteRequest) GetIds() string {
    return r.ids
}

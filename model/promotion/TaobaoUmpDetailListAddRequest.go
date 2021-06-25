package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
营销详情添加 APIRequest
taobao.ump.detail.list.add

批量添加营销活动。替代单条添加营销详情的的API。此接口适用针对某个营销活动的多档设置，会按顺序插入detail。若在整个事务过程中出现断点，会将已插入完成的detail_id返回，注意记录这些id，并将其删除，会对交易过程中的优惠产生影响。
*/
type TaobaoUmpDetailListAddRequest struct {
    model.Params

    // 营销活动id。
    actId   int64 

    // 营销详情的列表。此列表由detail的json字符串组成。最多插入为10个。
    details   string 

}

func NewTaobaoUmpDetailListAddRequest() *TaobaoUmpDetailListAddRequest{
    return &TaobaoUmpDetailListAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpDetailListAddRequest) GetApiMethodName() string {
    return "taobao.ump.detail.list.add"
}

func (r TaobaoUmpDetailListAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpDetailListAddRequest) SetActId(actId int64) error {
    r.actId = actId
    r.Set("act_id", actId)
    return nil
}

func (r TaobaoUmpDetailListAddRequest) GetActId() int64 {
    return r.actId
}

func (r *TaobaoUmpDetailListAddRequest) SetDetails(details string) error {
    r.details = details
    r.Set("details", details)
    return nil
}

func (r TaobaoUmpDetailListAddRequest) GetDetails() string {
    return r.details
}


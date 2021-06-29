package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
营销详情添加 API请求
taobao.ump.detail.list.add

批量添加营销活动。替代单条添加营销详情的的API。此接口适用针对某个营销活动的多档设置，会按顺序插入detail。若在整个事务过程中出现断点，会将已插入完成的detail_id返回，注意记录这些id，并将其删除，会对交易过程中的优惠产生影响。
*/
type TaobaoUmpDetailListAddRequest struct {
    model.Params
    // 营销活动id。
    _actId   int64
    // 营销详情的列表。此列表由detail的json字符串组成。最多插入为10个。
    _details   string
}

// 初始化TaobaoUmpDetailListAddRequest对象
func NewTaobaoUmpDetailListAddRequest() *TaobaoUmpDetailListAddRequest{
    return &TaobaoUmpDetailListAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpDetailListAddRequest) GetApiMethodName() string {
    return "taobao.ump.detail.list.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpDetailListAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActId Setter
// 营销活动id。
func (r *TaobaoUmpDetailListAddRequest) SetActId(_actId int64) error {
    r._actId = _actId
    r.Set("act_id", _actId)
    return nil
}

// ActId Getter
func (r TaobaoUmpDetailListAddRequest) GetActId() int64 {
    return r._actId
}
// Details Setter
// 营销详情的列表。此列表由detail的json字符串组成。最多插入为10个。
func (r *TaobaoUmpDetailListAddRequest) SetDetails(_details string) error {
    r._details = _details
    r.Set("details", _details)
    return nil
}

// Details Getter
func (r TaobaoUmpDetailListAddRequest) GetDetails() string {
    return r._details
}

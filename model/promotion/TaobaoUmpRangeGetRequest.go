package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动范围 API请求
taobao.ump.range.get

查询某个卖家所有参加或者不参加某项活动的物品
*/
type TaobaoUmpRangeGetRequest struct {
    model.Params
    // 活动id
    actId   int64
}

// 初始化TaobaoUmpRangeGetRequest对象
func NewTaobaoUmpRangeGetRequest() *TaobaoUmpRangeGetRequest{
    return &TaobaoUmpRangeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpRangeGetRequest) GetApiMethodName() string {
    return "taobao.ump.range.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpRangeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActId Setter
// 活动id
func (r *TaobaoUmpRangeGetRequest) SetActId(actId int64) error {
    r.actId = actId
    r.Set("act_id", actId)
    return nil
}

// ActId Getter
func (r TaobaoUmpRangeGetRequest) GetActId() int64 {
    return r.actId
}

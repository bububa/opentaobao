package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询营销活动 API请求
taobao.ump.activity.get

查询营销活动
*/
type TaobaoUmpActivityGetRequest struct {
    model.Params
    // 活动id
    actId   int64
}

// 初始化TaobaoUmpActivityGetRequest对象
func NewTaobaoUmpActivityGetRequest() *TaobaoUmpActivityGetRequest{
    return &TaobaoUmpActivityGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpActivityGetRequest) GetApiMethodName() string {
    return "taobao.ump.activity.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpActivityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActId Setter
// 活动id
func (r *TaobaoUmpActivityGetRequest) SetActId(actId int64) error {
    r.actId = actId
    r.Set("act_id", actId)
    return nil
}

// ActId Getter
func (r TaobaoUmpActivityGetRequest) GetActId() int64 {
    return r.actId
}

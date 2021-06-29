package degoperation

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
活动中奖记录 API请求
taobao.degoperation.show.top.records

活动中奖记录
*/
type TaobaoDegoperationShowTopRecordsRequest struct {
    model.Params
    // 活动后台配置
    _degAppKey   string
    // 活动后台配置
    _degEventKey   string
    // 返回数
    _topN   int64
}

// 初始化TaobaoDegoperationShowTopRecordsRequest对象
func NewTaobaoDegoperationShowTopRecordsRequest() *TaobaoDegoperationShowTopRecordsRequest{
    return &TaobaoDegoperationShowTopRecordsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDegoperationShowTopRecordsRequest) GetApiMethodName() string {
    return "taobao.degoperation.show.top.records"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDegoperationShowTopRecordsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DegAppKey Setter
// 活动后台配置
func (r *TaobaoDegoperationShowTopRecordsRequest) SetDegAppKey(_degAppKey string) error {
    r._degAppKey = _degAppKey
    r.Set("deg_app_key", _degAppKey)
    return nil
}

// DegAppKey Getter
func (r TaobaoDegoperationShowTopRecordsRequest) GetDegAppKey() string {
    return r._degAppKey
}
// DegEventKey Setter
// 活动后台配置
func (r *TaobaoDegoperationShowTopRecordsRequest) SetDegEventKey(_degEventKey string) error {
    r._degEventKey = _degEventKey
    r.Set("deg_event_key", _degEventKey)
    return nil
}

// DegEventKey Getter
func (r TaobaoDegoperationShowTopRecordsRequest) GetDegEventKey() string {
    return r._degEventKey
}
// TopN Setter
// 返回数
func (r *TaobaoDegoperationShowTopRecordsRequest) SetTopN(_topN int64) error {
    r._topN = _topN
    r.Set("top_n", _topN)
    return nil
}

// TopN Getter
func (r TaobaoDegoperationShowTopRecordsRequest) GetTopN() int64 {
    return r._topN
}

package degoperation

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
活动中奖记录 APIRequest
taobao.degoperation.show.top.records

活动中奖记录
*/
type TaobaoDegoperationShowTopRecordsRequest struct {
    model.Params

    // 活动后台配置
    degAppKey   string 

    // 活动后台配置
    degEventKey   string 

    // 返回数
    topN   int64 

}

func NewTaobaoDegoperationShowTopRecordsRequest() *TaobaoDegoperationShowTopRecordsRequest{
    return &TaobaoDegoperationShowTopRecordsRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDegoperationShowTopRecordsRequest) GetApiMethodName() string {
    return "taobao.degoperation.show.top.records"
}

func (r TaobaoDegoperationShowTopRecordsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDegoperationShowTopRecordsRequest) SetDegAppKey(degAppKey string) error {
    r.degAppKey = degAppKey
    r.Set("deg_app_key", degAppKey)
    return nil
}

func (r TaobaoDegoperationShowTopRecordsRequest) GetDegAppKey() string {
    return r.degAppKey
}

func (r *TaobaoDegoperationShowTopRecordsRequest) SetDegEventKey(degEventKey string) error {
    r.degEventKey = degEventKey
    r.Set("deg_event_key", degEventKey)
    return nil
}

func (r TaobaoDegoperationShowTopRecordsRequest) GetDegEventKey() string {
    return r.degEventKey
}

func (r *TaobaoDegoperationShowTopRecordsRequest) SetTopN(topN int64) error {
    r.topN = topN
    r.Set("top_n", topN)
    return nil
}

func (r TaobaoDegoperationShowTopRecordsRequest) GetTopN() int64 {
    return r.topN
}


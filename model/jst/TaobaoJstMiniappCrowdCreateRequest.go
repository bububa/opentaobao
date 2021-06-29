package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序活动创建 API请求
taobao.jst.miniapp.crowd.create

小程序活动创建
*/
type TaobaoJstMiniappCrowdCreateRequest struct {
    model.Params
    // 活动开始时间，开始时间和结束时间不能超过1个月
    _endDate   string
    // 活动描述
    _description   string
    // 活动开始时间
    _startDate   string
    // 活动名称
    _crowdName   string
}

// 初始化TaobaoJstMiniappCrowdCreateRequest对象
func NewTaobaoJstMiniappCrowdCreateRequest() *TaobaoJstMiniappCrowdCreateRequest{
    return &TaobaoJstMiniappCrowdCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstMiniappCrowdCreateRequest) GetApiMethodName() string {
    return "taobao.jst.miniapp.crowd.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstMiniappCrowdCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EndDate Setter
// 活动开始时间，开始时间和结束时间不能超过1个月
func (r *TaobaoJstMiniappCrowdCreateRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoJstMiniappCrowdCreateRequest) GetEndDate() string {
    return r._endDate
}
// Description Setter
// 活动描述
func (r *TaobaoJstMiniappCrowdCreateRequest) SetDescription(_description string) error {
    r._description = _description
    r.Set("description", _description)
    return nil
}

// Description Getter
func (r TaobaoJstMiniappCrowdCreateRequest) GetDescription() string {
    return r._description
}
// StartDate Setter
// 活动开始时间
func (r *TaobaoJstMiniappCrowdCreateRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoJstMiniappCrowdCreateRequest) GetStartDate() string {
    return r._startDate
}
// CrowdName Setter
// 活动名称
func (r *TaobaoJstMiniappCrowdCreateRequest) SetCrowdName(_crowdName string) error {
    r._crowdName = _crowdName
    r.Set("crowd_name", _crowdName)
    return nil
}

// CrowdName Getter
func (r TaobaoJstMiniappCrowdCreateRequest) GetCrowdName() string {
    return r._crowdName
}

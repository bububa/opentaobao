package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新物流进度 API请求
tmall.servicecenter.workcard.updatelogistics

提供给外部合作服务商的物流进度更改接口
*/
type TmallServicecenterWorkcardUpdatelogisticsRequest struct {
    model.Params
    // 工单号
    _workcardId   int64
    // 工单操作
    _action   string
    // 快递公司
    _expressCompany   string
    // 快递号
    _expressCode   string
}

// 初始化TmallServicecenterWorkcardUpdatelogisticsRequest对象
func NewTmallServicecenterWorkcardUpdatelogisticsRequest() *TmallServicecenterWorkcardUpdatelogisticsRequest{
    return &TmallServicecenterWorkcardUpdatelogisticsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardUpdatelogisticsRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.updatelogistics"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardUpdatelogisticsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkcardId Setter
// 工单号
func (r *TmallServicecenterWorkcardUpdatelogisticsRequest) SetWorkcardId(_workcardId int64) error {
    r._workcardId = _workcardId
    r.Set("workcard_id", _workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterWorkcardUpdatelogisticsRequest) GetWorkcardId() int64 {
    return r._workcardId
}
// Action Setter
// 工单操作
func (r *TmallServicecenterWorkcardUpdatelogisticsRequest) SetAction(_action string) error {
    r._action = _action
    r.Set("action", _action)
    return nil
}

// Action Getter
func (r TmallServicecenterWorkcardUpdatelogisticsRequest) GetAction() string {
    return r._action
}
// ExpressCompany Setter
// 快递公司
func (r *TmallServicecenterWorkcardUpdatelogisticsRequest) SetExpressCompany(_expressCompany string) error {
    r._expressCompany = _expressCompany
    r.Set("express_company", _expressCompany)
    return nil
}

// ExpressCompany Getter
func (r TmallServicecenterWorkcardUpdatelogisticsRequest) GetExpressCompany() string {
    return r._expressCompany
}
// ExpressCode Setter
// 快递号
func (r *TmallServicecenterWorkcardUpdatelogisticsRequest) SetExpressCode(_expressCode string) error {
    r._expressCode = _expressCode
    r.Set("express_code", _expressCode)
    return nil
}

// ExpressCode Getter
func (r TmallServicecenterWorkcardUpdatelogisticsRequest) GetExpressCode() string {
    return r._expressCode
}

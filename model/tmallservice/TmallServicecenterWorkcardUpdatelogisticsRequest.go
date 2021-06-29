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
    workcardId   int64
    // 工单操作
    action   string
    // 快递公司
    expressCompany   string
    // 快递号
    expressCode   string
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
func (r *TmallServicecenterWorkcardUpdatelogisticsRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterWorkcardUpdatelogisticsRequest) GetWorkcardId() int64 {
    return r.workcardId
}
// Action Setter
// 工单操作
func (r *TmallServicecenterWorkcardUpdatelogisticsRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

// Action Getter
func (r TmallServicecenterWorkcardUpdatelogisticsRequest) GetAction() string {
    return r.action
}
// ExpressCompany Setter
// 快递公司
func (r *TmallServicecenterWorkcardUpdatelogisticsRequest) SetExpressCompany(expressCompany string) error {
    r.expressCompany = expressCompany
    r.Set("express_company", expressCompany)
    return nil
}

// ExpressCompany Getter
func (r TmallServicecenterWorkcardUpdatelogisticsRequest) GetExpressCompany() string {
    return r.expressCompany
}
// ExpressCode Setter
// 快递号
func (r *TmallServicecenterWorkcardUpdatelogisticsRequest) SetExpressCode(expressCode string) error {
    r.expressCode = expressCode
    r.Set("express_code", expressCode)
    return nil
}

// ExpressCode Getter
func (r TmallServicecenterWorkcardUpdatelogisticsRequest) GetExpressCode() string {
    return r.expressCode
}

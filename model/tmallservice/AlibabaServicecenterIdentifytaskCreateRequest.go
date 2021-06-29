package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建核销单 API请求
alibaba.servicecenter.identifytask.create

创建核销单
*/
type AlibabaServicecenterIdentifytaskCreateRequest struct {
    model.Params
    // 工单集合
    workcardIds   []int64
    // 核销单外部标识，服务商保证唯一。如果创建核销单时传入重复的outer_id，系统直接会返回服务商下该outer_id对应的核销单，不会重新创建新核销单。
    outerId   string
}

// 初始化AlibabaServicecenterIdentifytaskCreateRequest对象
func NewAlibabaServicecenterIdentifytaskCreateRequest() *AlibabaServicecenterIdentifytaskCreateRequest{
    return &AlibabaServicecenterIdentifytaskCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterIdentifytaskCreateRequest) GetApiMethodName() string {
    return "alibaba.servicecenter.identifytask.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterIdentifytaskCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkcardIds Setter
// 工单集合
func (r *AlibabaServicecenterIdentifytaskCreateRequest) SetWorkcardIds(workcardIds []int64) error {
    r.workcardIds = workcardIds
    r.Set("workcard_ids", workcardIds)
    return nil
}

// WorkcardIds Getter
func (r AlibabaServicecenterIdentifytaskCreateRequest) GetWorkcardIds() []int64 {
    return r.workcardIds
}
// OuterId Setter
// 核销单外部标识，服务商保证唯一。如果创建核销单时传入重复的outer_id，系统直接会返回服务商下该outer_id对应的核销单，不会重新创建新核销单。
func (r *AlibabaServicecenterIdentifytaskCreateRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r AlibabaServicecenterIdentifytaskCreateRequest) GetOuterId() string {
    return r.outerId
}

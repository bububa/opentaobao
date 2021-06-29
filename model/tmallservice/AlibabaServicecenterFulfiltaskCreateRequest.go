package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
合单生成核销单 API请求
alibaba.servicecenter.fulfiltask.create

服务对工单进行合单，合单的结果是生成核销单
*/
type AlibabaServicecenterFulfiltaskCreateRequest struct {
    model.Params
    // 工单id列表
    _workcardIds   []int64
    // 外部单号
    _outerId   string
}

// 初始化AlibabaServicecenterFulfiltaskCreateRequest对象
func NewAlibabaServicecenterFulfiltaskCreateRequest() *AlibabaServicecenterFulfiltaskCreateRequest{
    return &AlibabaServicecenterFulfiltaskCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterFulfiltaskCreateRequest) GetApiMethodName() string {
    return "alibaba.servicecenter.fulfiltask.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterFulfiltaskCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkcardIds Setter
// 工单id列表
func (r *AlibabaServicecenterFulfiltaskCreateRequest) SetWorkcardIds(_workcardIds []int64) error {
    r._workcardIds = _workcardIds
    r.Set("workcard_ids", _workcardIds)
    return nil
}

// WorkcardIds Getter
func (r AlibabaServicecenterFulfiltaskCreateRequest) GetWorkcardIds() []int64 {
    return r._workcardIds
}
// OuterId Setter
// 外部单号
func (r *AlibabaServicecenterFulfiltaskCreateRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r AlibabaServicecenterFulfiltaskCreateRequest) GetOuterId() string {
    return r._outerId
}

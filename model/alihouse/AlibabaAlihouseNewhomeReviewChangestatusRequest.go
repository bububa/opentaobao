package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
楼盘测评草稿状态同步 API请求
alibaba.alihouse.newhome.review.changestatus

楼盘测评草稿状态更新
*/
type AlibabaAlihouseNewhomeReviewChangestatusRequest struct {
    model.Params
    // 外部测评id
    _outerId   string
    // 0 失效 1 有效
    _status   int64
}

// 初始化AlibabaAlihouseNewhomeReviewChangestatusRequest对象
func NewAlibabaAlihouseNewhomeReviewChangestatusRequest() *AlibabaAlihouseNewhomeReviewChangestatusRequest{
    return &AlibabaAlihouseNewhomeReviewChangestatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeReviewChangestatusRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.review.changestatus"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeReviewChangestatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterId Setter
// 外部测评id
func (r *AlibabaAlihouseNewhomeReviewChangestatusRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r AlibabaAlihouseNewhomeReviewChangestatusRequest) GetOuterId() string {
    return r._outerId
}
// Status Setter
// 0 失效 1 有效
func (r *AlibabaAlihouseNewhomeReviewChangestatusRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaAlihouseNewhomeReviewChangestatusRequest) GetStatus() int64 {
    return r._status
}

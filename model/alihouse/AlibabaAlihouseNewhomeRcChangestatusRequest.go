package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图文草稿状态更新 API请求
alibaba.alihouse.newhome.rc.changestatus

图文草稿状态更新
*/
type AlibabaAlihouseNewhomeRcChangestatusRequest struct {
    model.Params
    // 外部图文id
    _outerId   string
    // 0 失效 1 有效
    _status   int64
}

// 初始化AlibabaAlihouseNewhomeRcChangestatusRequest对象
func NewAlibabaAlihouseNewhomeRcChangestatusRequest() *AlibabaAlihouseNewhomeRcChangestatusRequest{
    return &AlibabaAlihouseNewhomeRcChangestatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeRcChangestatusRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.rc.changestatus"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeRcChangestatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterId Setter
// 外部图文id
func (r *AlibabaAlihouseNewhomeRcChangestatusRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r AlibabaAlihouseNewhomeRcChangestatusRequest) GetOuterId() string {
    return r._outerId
}
// Status Setter
// 0 失效 1 有效
func (r *AlibabaAlihouseNewhomeRcChangestatusRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaAlihouseNewhomeRcChangestatusRequest) GetStatus() int64 {
    return r._status
}

package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
视频草稿状态更新 API请求
alibaba.alihouse.newhome.video.changestatus

视频草稿状态更新
*/
type AlibabaAlihouseNewhomeVideoChangestatusRequest struct {
    model.Params
    // 外部视频id
    _outerId   string
    // 0 失效 1 有效
    _status   int64
}

// 初始化AlibabaAlihouseNewhomeVideoChangestatusRequest对象
func NewAlibabaAlihouseNewhomeVideoChangestatusRequest() *AlibabaAlihouseNewhomeVideoChangestatusRequest{
    return &AlibabaAlihouseNewhomeVideoChangestatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeVideoChangestatusRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.video.changestatus"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeVideoChangestatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterId Setter
// 外部视频id
func (r *AlibabaAlihouseNewhomeVideoChangestatusRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r AlibabaAlihouseNewhomeVideoChangestatusRequest) GetOuterId() string {
    return r._outerId
}
// Status Setter
// 0 失效 1 有效
func (r *AlibabaAlihouseNewhomeVideoChangestatusRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaAlihouseNewhomeVideoChangestatusRequest) GetStatus() int64 {
    return r._status
}

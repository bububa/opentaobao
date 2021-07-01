package youkuott

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
dvb工单反馈 API请求
youku.ott.dvb.workorder.feedback

dvb工单处理结果反馈
*/
type YoukuOttDvbWorkorderFeedbackAPIRequest struct {
    model.Params
    // 工单id
    _workorderId   int64
    // 反馈内容
    _content   string
    // 操作发生时间（时间戳：毫秒）
    _occureTime   int64
}

// 初始化YoukuOttDvbWorkorderFeedbackAPIRequest对象
func NewYoukuOttDvbWorkorderFeedbackRequest() *YoukuOttDvbWorkorderFeedbackAPIRequest{
    return &YoukuOttDvbWorkorderFeedbackAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuOttDvbWorkorderFeedbackAPIRequest) GetApiMethodName() string {
    return "youku.ott.dvb.workorder.feedback"
}

// IRequest interface 方法, 获取API参数
func (r YoukuOttDvbWorkorderFeedbackAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkorderId Setter
// 工单id
func (r *YoukuOttDvbWorkorderFeedbackAPIRequest) SetWorkorderId(_workorderId int64) error {
    r._workorderId = _workorderId
    r.Set("workorder_id", _workorderId)
    return nil
}

// WorkorderId Getter
func (r YoukuOttDvbWorkorderFeedbackAPIRequest) GetWorkorderId() int64 {
    return r._workorderId
}
// Content Setter
// 反馈内容
func (r *YoukuOttDvbWorkorderFeedbackAPIRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r YoukuOttDvbWorkorderFeedbackAPIRequest) GetContent() string {
    return r._content
}
// OccureTime Setter
// 操作发生时间（时间戳：毫秒）
func (r *YoukuOttDvbWorkorderFeedbackAPIRequest) SetOccureTime(_occureTime int64) error {
    r._occureTime = _occureTime
    r.Set("occure_time", _occureTime)
    return nil
}

// OccureTime Getter
func (r YoukuOttDvbWorkorderFeedbackAPIRequest) GetOccureTime() int64 {
    return r._occureTime
}

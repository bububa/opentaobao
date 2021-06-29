package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广淘小铺isv 活动到微淘feed API请求
taobao.weitao.feed.synchronize

推广淘小铺isv 活动到微淘feed
*/
type TaobaoWeitaoFeedSynchronizeRequest struct {
    model.Params
    // 活动id
    _bizId   int64
    // feed封面图片url
    _coverPath   string
    // feed点击跳转的活动地址
    _detailUrl   string
    // feed展示结束时间
    _endTime   int64
    // feed展示开始时间
    _startTime   int64
    // feed描述
    _summary   string
    // feed标题
    _title   string
}

// 初始化TaobaoWeitaoFeedSynchronizeRequest对象
func NewTaobaoWeitaoFeedSynchronizeRequest() *TaobaoWeitaoFeedSynchronizeRequest{
    return &TaobaoWeitaoFeedSynchronizeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWeitaoFeedSynchronizeRequest) GetApiMethodName() string {
    return "taobao.weitao.feed.synchronize"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWeitaoFeedSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizId Setter
// 活动id
func (r *TaobaoWeitaoFeedSynchronizeRequest) SetBizId(_bizId int64) error {
    r._bizId = _bizId
    r.Set("biz_id", _bizId)
    return nil
}

// BizId Getter
func (r TaobaoWeitaoFeedSynchronizeRequest) GetBizId() int64 {
    return r._bizId
}
// CoverPath Setter
// feed封面图片url
func (r *TaobaoWeitaoFeedSynchronizeRequest) SetCoverPath(_coverPath string) error {
    r._coverPath = _coverPath
    r.Set("cover_path", _coverPath)
    return nil
}

// CoverPath Getter
func (r TaobaoWeitaoFeedSynchronizeRequest) GetCoverPath() string {
    return r._coverPath
}
// DetailUrl Setter
// feed点击跳转的活动地址
func (r *TaobaoWeitaoFeedSynchronizeRequest) SetDetailUrl(_detailUrl string) error {
    r._detailUrl = _detailUrl
    r.Set("detail_url", _detailUrl)
    return nil
}

// DetailUrl Getter
func (r TaobaoWeitaoFeedSynchronizeRequest) GetDetailUrl() string {
    return r._detailUrl
}
// EndTime Setter
// feed展示结束时间
func (r *TaobaoWeitaoFeedSynchronizeRequest) SetEndTime(_endTime int64) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoWeitaoFeedSynchronizeRequest) GetEndTime() int64 {
    return r._endTime
}
// StartTime Setter
// feed展示开始时间
func (r *TaobaoWeitaoFeedSynchronizeRequest) SetStartTime(_startTime int64) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoWeitaoFeedSynchronizeRequest) GetStartTime() int64 {
    return r._startTime
}
// Summary Setter
// feed描述
func (r *TaobaoWeitaoFeedSynchronizeRequest) SetSummary(_summary string) error {
    r._summary = _summary
    r.Set("summary", _summary)
    return nil
}

// Summary Getter
func (r TaobaoWeitaoFeedSynchronizeRequest) GetSummary() string {
    return r._summary
}
// Title Setter
// feed标题
func (r *TaobaoWeitaoFeedSynchronizeRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r TaobaoWeitaoFeedSynchronizeRequest) GetTitle() string {
    return r._title
}

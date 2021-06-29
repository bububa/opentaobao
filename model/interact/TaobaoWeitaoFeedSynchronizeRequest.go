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
    bizId   int64
    // feed封面图片url
    coverPath   string
    // feed点击跳转的活动地址
    detailUrl   string
    // feed展示结束时间
    endTime   int64
    // feed展示开始时间
    startTime   int64
    // feed描述
    summary   string
    // feed标题
    title   string
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
func (r *TaobaoWeitaoFeedSynchronizeRequest) SetBizId(bizId int64) error {
    r.bizId = bizId
    r.Set("biz_id", bizId)
    return nil
}

// BizId Getter
func (r TaobaoWeitaoFeedSynchronizeRequest) GetBizId() int64 {
    return r.bizId
}
// CoverPath Setter
// feed封面图片url
func (r *TaobaoWeitaoFeedSynchronizeRequest) SetCoverPath(coverPath string) error {
    r.coverPath = coverPath
    r.Set("cover_path", coverPath)
    return nil
}

// CoverPath Getter
func (r TaobaoWeitaoFeedSynchronizeRequest) GetCoverPath() string {
    return r.coverPath
}
// DetailUrl Setter
// feed点击跳转的活动地址
func (r *TaobaoWeitaoFeedSynchronizeRequest) SetDetailUrl(detailUrl string) error {
    r.detailUrl = detailUrl
    r.Set("detail_url", detailUrl)
    return nil
}

// DetailUrl Getter
func (r TaobaoWeitaoFeedSynchronizeRequest) GetDetailUrl() string {
    return r.detailUrl
}
// EndTime Setter
// feed展示结束时间
func (r *TaobaoWeitaoFeedSynchronizeRequest) SetEndTime(endTime int64) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoWeitaoFeedSynchronizeRequest) GetEndTime() int64 {
    return r.endTime
}
// StartTime Setter
// feed展示开始时间
func (r *TaobaoWeitaoFeedSynchronizeRequest) SetStartTime(startTime int64) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoWeitaoFeedSynchronizeRequest) GetStartTime() int64 {
    return r.startTime
}
// Summary Setter
// feed描述
func (r *TaobaoWeitaoFeedSynchronizeRequest) SetSummary(summary string) error {
    r.summary = summary
    r.Set("summary", summary)
    return nil
}

// Summary Getter
func (r TaobaoWeitaoFeedSynchronizeRequest) GetSummary() string {
    return r.summary
}
// Title Setter
// feed标题
func (r *TaobaoWeitaoFeedSynchronizeRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TaobaoWeitaoFeedSynchronizeRequest) GetTitle() string {
    return r.title
}

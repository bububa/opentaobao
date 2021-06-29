package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广淘小铺isv 活动到微淘feed API请求
taobao.weitao.feed.synchronize.new

推广微淘互动应用活动到微淘
*/
type TaobaoWeitaoFeedSynchronizeNewRequest struct {
    model.Params
    // 广播类型：粉丝猜价格461、投票有礼462、粉丝抢红包463、盖楼有礼464、加购有礼465
    feedType   int64
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
    // 宝贝列表，用于card展示，0~2个宝贝ID
    itemIds   []string
    // 活动ID
    sbizId   string
}

// 初始化TaobaoWeitaoFeedSynchronizeNewRequest对象
func NewTaobaoWeitaoFeedSynchronizeNewRequest() *TaobaoWeitaoFeedSynchronizeNewRequest{
    return &TaobaoWeitaoFeedSynchronizeNewRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetApiMethodName() string {
    return "taobao.weitao.feed.synchronize.new"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FeedType Setter
// 广播类型：粉丝猜价格461、投票有礼462、粉丝抢红包463、盖楼有礼464、加购有礼465
func (r *TaobaoWeitaoFeedSynchronizeNewRequest) SetFeedType(feedType int64) error {
    r.feedType = feedType
    r.Set("feed_type", feedType)
    return nil
}

// FeedType Getter
func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetFeedType() int64 {
    return r.feedType
}
// DetailUrl Setter
// feed点击跳转的活动地址
func (r *TaobaoWeitaoFeedSynchronizeNewRequest) SetDetailUrl(detailUrl string) error {
    r.detailUrl = detailUrl
    r.Set("detail_url", detailUrl)
    return nil
}

// DetailUrl Getter
func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetDetailUrl() string {
    return r.detailUrl
}
// EndTime Setter
// feed展示结束时间
func (r *TaobaoWeitaoFeedSynchronizeNewRequest) SetEndTime(endTime int64) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetEndTime() int64 {
    return r.endTime
}
// StartTime Setter
// feed展示开始时间
func (r *TaobaoWeitaoFeedSynchronizeNewRequest) SetStartTime(startTime int64) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetStartTime() int64 {
    return r.startTime
}
// Summary Setter
// feed描述
func (r *TaobaoWeitaoFeedSynchronizeNewRequest) SetSummary(summary string) error {
    r.summary = summary
    r.Set("summary", summary)
    return nil
}

// Summary Getter
func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetSummary() string {
    return r.summary
}
// Title Setter
// feed标题
func (r *TaobaoWeitaoFeedSynchronizeNewRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetTitle() string {
    return r.title
}
// ItemIds Setter
// 宝贝列表，用于card展示，0~2个宝贝ID
func (r *TaobaoWeitaoFeedSynchronizeNewRequest) SetItemIds(itemIds []string) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetItemIds() []string {
    return r.itemIds
}
// SbizId Setter
// 活动ID
func (r *TaobaoWeitaoFeedSynchronizeNewRequest) SetSbizId(sbizId string) error {
    r.sbizId = sbizId
    r.Set("sbiz_id", sbizId)
    return nil
}

// SbizId Getter
func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetSbizId() string {
    return r.sbizId
}

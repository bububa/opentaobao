package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广淘小铺isv 活动到微淘feed APIRequest
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

func NewTaobaoWeitaoFeedSynchronizeNewRequest() *TaobaoWeitaoFeedSynchronizeNewRequest{
    return &TaobaoWeitaoFeedSynchronizeNewRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetApiMethodName() string {
    return "taobao.weitao.feed.synchronize.new"
}

func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWeitaoFeedSynchronizeNewRequest) SetFeedType(feedType int64) error {
    r.feedType = feedType
    r.Set("feed_type", feedType)
    return nil
}

func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetFeedType() int64 {
    return r.feedType
}

func (r *TaobaoWeitaoFeedSynchronizeNewRequest) SetDetailUrl(detailUrl string) error {
    r.detailUrl = detailUrl
    r.Set("detail_url", detailUrl)
    return nil
}

func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetDetailUrl() string {
    return r.detailUrl
}

func (r *TaobaoWeitaoFeedSynchronizeNewRequest) SetEndTime(endTime int64) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetEndTime() int64 {
    return r.endTime
}

func (r *TaobaoWeitaoFeedSynchronizeNewRequest) SetStartTime(startTime int64) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetStartTime() int64 {
    return r.startTime
}

func (r *TaobaoWeitaoFeedSynchronizeNewRequest) SetSummary(summary string) error {
    r.summary = summary
    r.Set("summary", summary)
    return nil
}

func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetSummary() string {
    return r.summary
}

func (r *TaobaoWeitaoFeedSynchronizeNewRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetTitle() string {
    return r.title
}

func (r *TaobaoWeitaoFeedSynchronizeNewRequest) SetItemIds(itemIds []string) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetItemIds() []string {
    return r.itemIds
}

func (r *TaobaoWeitaoFeedSynchronizeNewRequest) SetSbizId(sbizId string) error {
    r.sbizId = sbizId
    r.Set("sbiz_id", sbizId)
    return nil
}

func (r TaobaoWeitaoFeedSynchronizeNewRequest) GetSbizId() string {
    return r.sbizId
}


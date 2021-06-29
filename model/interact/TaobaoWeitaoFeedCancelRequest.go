package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消广播在timeline、广场中展示 API请求
taobao.weitao.feed.cancel

取消广播在timeline和广场中的展示。
*/
type TaobaoWeitaoFeedCancelRequest struct {
    model.Params
    // 广播id
    feedId   int64
    // 三方活动ID
    bizId   string
    // 是否彻底删除（店铺动态不可见，等同卖家广播后台删除），默认false
    delete   bool
}

// 初始化TaobaoWeitaoFeedCancelRequest对象
func NewTaobaoWeitaoFeedCancelRequest() *TaobaoWeitaoFeedCancelRequest{
    return &TaobaoWeitaoFeedCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWeitaoFeedCancelRequest) GetApiMethodName() string {
    return "taobao.weitao.feed.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWeitaoFeedCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FeedId Setter
// 广播id
func (r *TaobaoWeitaoFeedCancelRequest) SetFeedId(feedId int64) error {
    r.feedId = feedId
    r.Set("feed_id", feedId)
    return nil
}

// FeedId Getter
func (r TaobaoWeitaoFeedCancelRequest) GetFeedId() int64 {
    return r.feedId
}
// BizId Setter
// 三方活动ID
func (r *TaobaoWeitaoFeedCancelRequest) SetBizId(bizId string) error {
    r.bizId = bizId
    r.Set("biz_id", bizId)
    return nil
}

// BizId Getter
func (r TaobaoWeitaoFeedCancelRequest) GetBizId() string {
    return r.bizId
}
// Delete Setter
// 是否彻底删除（店铺动态不可见，等同卖家广播后台删除），默认false
func (r *TaobaoWeitaoFeedCancelRequest) SetDelete(delete bool) error {
    r.delete = delete
    r.Set("delete", delete)
    return nil
}

// Delete Getter
func (r TaobaoWeitaoFeedCancelRequest) GetDelete() bool {
    return r.delete
}

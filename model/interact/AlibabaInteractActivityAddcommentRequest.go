package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
微淘评论接口 API请求
alibaba.interact.activity.addcomment

发表评论，并返回楼层
*/
type AlibabaInteractActivityAddcommentRequest struct {
    model.Params
    // 该字段为评论内容
    content   string
    // 评论feedid
    feedId   int64
    // 发评论的业务id
    bizId   string
}

// 初始化AlibabaInteractActivityAddcommentRequest对象
func NewAlibabaInteractActivityAddcommentRequest() *AlibabaInteractActivityAddcommentRequest{
    return &AlibabaInteractActivityAddcommentRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractActivityAddcommentRequest) GetApiMethodName() string {
    return "alibaba.interact.activity.addcomment"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractActivityAddcommentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Content Setter
// 该字段为评论内容
func (r *AlibabaInteractActivityAddcommentRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r AlibabaInteractActivityAddcommentRequest) GetContent() string {
    return r.content
}
// FeedId Setter
// 评论feedid
func (r *AlibabaInteractActivityAddcommentRequest) SetFeedId(feedId int64) error {
    r.feedId = feedId
    r.Set("feed_id", feedId)
    return nil
}

// FeedId Getter
func (r AlibabaInteractActivityAddcommentRequest) GetFeedId() int64 {
    return r.feedId
}
// BizId Setter
// 发评论的业务id
func (r *AlibabaInteractActivityAddcommentRequest) SetBizId(bizId string) error {
    r.bizId = bizId
    r.Set("biz_id", bizId)
    return nil
}

// BizId Getter
func (r AlibabaInteractActivityAddcommentRequest) GetBizId() string {
    return r.bizId
}

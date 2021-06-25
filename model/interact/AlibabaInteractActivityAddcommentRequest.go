package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
微淘评论接口 APIRequest
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

func NewAlibabaInteractActivityAddcommentRequest() *AlibabaInteractActivityAddcommentRequest{
    return &AlibabaInteractActivityAddcommentRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractActivityAddcommentRequest) GetApiMethodName() string {
    return "alibaba.interact.activity.addcomment"
}

func (r AlibabaInteractActivityAddcommentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractActivityAddcommentRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r AlibabaInteractActivityAddcommentRequest) GetContent() string {
    return r.content
}

func (r *AlibabaInteractActivityAddcommentRequest) SetFeedId(feedId int64) error {
    r.feedId = feedId
    r.Set("feed_id", feedId)
    return nil
}

func (r AlibabaInteractActivityAddcommentRequest) GetFeedId() int64 {
    return r.feedId
}

func (r *AlibabaInteractActivityAddcommentRequest) SetBizId(bizId string) error {
    r.bizId = bizId
    r.Set("biz_id", bizId)
    return nil
}

func (r AlibabaInteractActivityAddcommentRequest) GetBizId() string {
    return r.bizId
}


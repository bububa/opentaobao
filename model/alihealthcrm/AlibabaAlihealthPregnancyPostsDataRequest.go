package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发回帖子信息同步 APIRequest
alibaba.alihealth.pregnancy.posts.data

发回帖子信息同步
*/
type AlibabaAlihealthPregnancyPostsDataRequest struct {
    model.Params

    // 用户id
    userId   int64 

    // 事件类型 0发帖 1回帖
    eventType   int64 

    // 主贴id
    mainId   int64 

    // 回帖id
    replyId   int64 

    // 标题
    title   string 

    // 内容
    content   string 

    // 图片url
    picUrl   string 

    // 发帖时间
    date   int64 

}

func NewAlibabaAlihealthPregnancyPostsDataRequest() *AlibabaAlihealthPregnancyPostsDataRequest{
    return &AlibabaAlihealthPregnancyPostsDataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthPregnancyPostsDataRequest) GetApiMethodName() string {
    return "alibaba.alihealth.pregnancy.posts.data"
}

func (r AlibabaAlihealthPregnancyPostsDataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthPregnancyPostsDataRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaAlihealthPregnancyPostsDataRequest) GetUserId() int64 {
    return r.userId
}

func (r *AlibabaAlihealthPregnancyPostsDataRequest) SetEventType(eventType int64) error {
    r.eventType = eventType
    r.Set("event_type", eventType)
    return nil
}

func (r AlibabaAlihealthPregnancyPostsDataRequest) GetEventType() int64 {
    return r.eventType
}

func (r *AlibabaAlihealthPregnancyPostsDataRequest) SetMainId(mainId int64) error {
    r.mainId = mainId
    r.Set("main_id", mainId)
    return nil
}

func (r AlibabaAlihealthPregnancyPostsDataRequest) GetMainId() int64 {
    return r.mainId
}

func (r *AlibabaAlihealthPregnancyPostsDataRequest) SetReplyId(replyId int64) error {
    r.replyId = replyId
    r.Set("reply_id", replyId)
    return nil
}

func (r AlibabaAlihealthPregnancyPostsDataRequest) GetReplyId() int64 {
    return r.replyId
}

func (r *AlibabaAlihealthPregnancyPostsDataRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r AlibabaAlihealthPregnancyPostsDataRequest) GetTitle() string {
    return r.title
}

func (r *AlibabaAlihealthPregnancyPostsDataRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r AlibabaAlihealthPregnancyPostsDataRequest) GetContent() string {
    return r.content
}

func (r *AlibabaAlihealthPregnancyPostsDataRequest) SetPicUrl(picUrl string) error {
    r.picUrl = picUrl
    r.Set("pic_url", picUrl)
    return nil
}

func (r AlibabaAlihealthPregnancyPostsDataRequest) GetPicUrl() string {
    return r.picUrl
}

func (r *AlibabaAlihealthPregnancyPostsDataRequest) SetDate(date int64) error {
    r.date = date
    r.Set("date", date)
    return nil
}

func (r AlibabaAlihealthPregnancyPostsDataRequest) GetDate() int64 {
    return r.date
}


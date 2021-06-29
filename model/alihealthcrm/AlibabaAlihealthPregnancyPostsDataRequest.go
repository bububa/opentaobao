package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发回帖子信息同步 API请求
alibaba.alihealth.pregnancy.posts.data

发回帖子信息同步
*/
type AlibabaAlihealthPregnancyPostsDataRequest struct {
    model.Params
    // 用户id
    _userId   int64
    // 事件类型 0发帖 1回帖
    _eventType   int64
    // 主贴id
    _mainId   int64
    // 回帖id
    _replyId   int64
    // 标题
    _title   string
    // 内容
    _content   string
    // 图片url
    _picUrl   string
    // 发帖时间
    _date   int64
}

// 初始化AlibabaAlihealthPregnancyPostsDataRequest对象
func NewAlibabaAlihealthPregnancyPostsDataRequest() *AlibabaAlihealthPregnancyPostsDataRequest{
    return &AlibabaAlihealthPregnancyPostsDataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPregnancyPostsDataRequest) GetApiMethodName() string {
    return "alibaba.alihealth.pregnancy.posts.data"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPregnancyPostsDataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户id
func (r *AlibabaAlihealthPregnancyPostsDataRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaAlihealthPregnancyPostsDataRequest) GetUserId() int64 {
    return r._userId
}
// EventType Setter
// 事件类型 0发帖 1回帖
func (r *AlibabaAlihealthPregnancyPostsDataRequest) SetEventType(_eventType int64) error {
    r._eventType = _eventType
    r.Set("event_type", _eventType)
    return nil
}

// EventType Getter
func (r AlibabaAlihealthPregnancyPostsDataRequest) GetEventType() int64 {
    return r._eventType
}
// MainId Setter
// 主贴id
func (r *AlibabaAlihealthPregnancyPostsDataRequest) SetMainId(_mainId int64) error {
    r._mainId = _mainId
    r.Set("main_id", _mainId)
    return nil
}

// MainId Getter
func (r AlibabaAlihealthPregnancyPostsDataRequest) GetMainId() int64 {
    return r._mainId
}
// ReplyId Setter
// 回帖id
func (r *AlibabaAlihealthPregnancyPostsDataRequest) SetReplyId(_replyId int64) error {
    r._replyId = _replyId
    r.Set("reply_id", _replyId)
    return nil
}

// ReplyId Getter
func (r AlibabaAlihealthPregnancyPostsDataRequest) GetReplyId() int64 {
    return r._replyId
}
// Title Setter
// 标题
func (r *AlibabaAlihealthPregnancyPostsDataRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r AlibabaAlihealthPregnancyPostsDataRequest) GetTitle() string {
    return r._title
}
// Content Setter
// 内容
func (r *AlibabaAlihealthPregnancyPostsDataRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r AlibabaAlihealthPregnancyPostsDataRequest) GetContent() string {
    return r._content
}
// PicUrl Setter
// 图片url
func (r *AlibabaAlihealthPregnancyPostsDataRequest) SetPicUrl(_picUrl string) error {
    r._picUrl = _picUrl
    r.Set("pic_url", _picUrl)
    return nil
}

// PicUrl Getter
func (r AlibabaAlihealthPregnancyPostsDataRequest) GetPicUrl() string {
    return r._picUrl
}
// Date Setter
// 发帖时间
func (r *AlibabaAlihealthPregnancyPostsDataRequest) SetDate(_date int64) error {
    r._date = _date
    r.Set("date", _date)
    return nil
}

// Date Getter
func (r AlibabaAlihealthPregnancyPostsDataRequest) GetDate() int64 {
    return r._date
}

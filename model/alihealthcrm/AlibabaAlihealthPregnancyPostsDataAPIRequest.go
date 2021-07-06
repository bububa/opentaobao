package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPregnancyPostsDataAPIRequest 发回帖子信息同步 API请求
// alibaba.alihealth.pregnancy.posts.data
//
// 发回帖子信息同步
type AlibabaAlihealthPregnancyPostsDataAPIRequest struct {
	model.Params
	// 标题
	_title string
	// 内容
	_content string
	// 图片url
	_picUrl string
	// 用户id
	_userId int64
	// 事件类型 0发帖 1回帖
	_eventType int64
	// 主贴id
	_mainId int64
	// 回帖id
	_replyId int64
	// 发帖时间
	_date int64
}

// NewAlibabaAlihealthPregnancyPostsDataRequest 初始化AlibabaAlihealthPregnancyPostsDataAPIRequest对象
func NewAlibabaAlihealthPregnancyPostsDataRequest() *AlibabaAlihealthPregnancyPostsDataAPIRequest {
	return &AlibabaAlihealthPregnancyPostsDataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPregnancyPostsDataAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pregnancy.posts.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPregnancyPostsDataAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTitle is Title Setter
// 标题
func (r *AlibabaAlihealthPregnancyPostsDataAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r AlibabaAlihealthPregnancyPostsDataAPIRequest) GetTitle() string {
	return r._title
}

// SetContent is Content Setter
// 内容
func (r *AlibabaAlihealthPregnancyPostsDataAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabaAlihealthPregnancyPostsDataAPIRequest) GetContent() string {
	return r._content
}

// SetPicUrl is PicUrl Setter
// 图片url
func (r *AlibabaAlihealthPregnancyPostsDataAPIRequest) SetPicUrl(_picUrl string) error {
	r._picUrl = _picUrl
	r.Set("pic_url", _picUrl)
	return nil
}

// GetPicUrl PicUrl Getter
func (r AlibabaAlihealthPregnancyPostsDataAPIRequest) GetPicUrl() string {
	return r._picUrl
}

// SetUserId is UserId Setter
// 用户id
func (r *AlibabaAlihealthPregnancyPostsDataAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaAlihealthPregnancyPostsDataAPIRequest) GetUserId() int64 {
	return r._userId
}

// SetEventType is EventType Setter
// 事件类型 0发帖 1回帖
func (r *AlibabaAlihealthPregnancyPostsDataAPIRequest) SetEventType(_eventType int64) error {
	r._eventType = _eventType
	r.Set("event_type", _eventType)
	return nil
}

// GetEventType EventType Getter
func (r AlibabaAlihealthPregnancyPostsDataAPIRequest) GetEventType() int64 {
	return r._eventType
}

// SetMainId is MainId Setter
// 主贴id
func (r *AlibabaAlihealthPregnancyPostsDataAPIRequest) SetMainId(_mainId int64) error {
	r._mainId = _mainId
	r.Set("main_id", _mainId)
	return nil
}

// GetMainId MainId Getter
func (r AlibabaAlihealthPregnancyPostsDataAPIRequest) GetMainId() int64 {
	return r._mainId
}

// SetReplyId is ReplyId Setter
// 回帖id
func (r *AlibabaAlihealthPregnancyPostsDataAPIRequest) SetReplyId(_replyId int64) error {
	r._replyId = _replyId
	r.Set("reply_id", _replyId)
	return nil
}

// GetReplyId ReplyId Getter
func (r AlibabaAlihealthPregnancyPostsDataAPIRequest) GetReplyId() int64 {
	return r._replyId
}

// SetDate is Date Setter
// 发帖时间
func (r *AlibabaAlihealthPregnancyPostsDataAPIRequest) SetDate(_date int64) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r AlibabaAlihealthPregnancyPostsDataAPIRequest) GetDate() int64 {
	return r._date
}

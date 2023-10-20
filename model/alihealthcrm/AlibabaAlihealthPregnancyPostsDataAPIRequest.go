package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpregnancypostsdataAPIRequest 发回帖子信息同步 API请求
// alibaba.alihealth.pregnancy.posts.data
//
// 发回帖子信息同步
type AlibabaalihealthpregnancypostsdataAPIRequest struct {
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

// NewAlibabaalihealthpregnancypostsdataRequest 初始化AlibabaalihealthpregnancypostsdataAPIRequest对象
func NewAlibabaalihealthpregnancypostsdataRequest() *AlibabaalihealthpregnancypostsdataAPIRequest {
	return &AlibabaalihealthpregnancypostsdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthpregnancypostsdataAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pregnancy.posts.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthpregnancypostsdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthpregnancypostsdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTitle is Title Setter
// 标题
func (r *AlibabaalihealthpregnancypostsdataAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r AlibabaalihealthpregnancypostsdataAPIRequest) GetTitle() string {
	return r._title
}

// SetContent is Content Setter
// 内容
func (r *AlibabaalihealthpregnancypostsdataAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabaalihealthpregnancypostsdataAPIRequest) GetContent() string {
	return r._content
}

// SetPicUrl is PicUrl Setter
// 图片url
func (r *AlibabaalihealthpregnancypostsdataAPIRequest) SetPicUrl(_picUrl string) error {
	r._picUrl = _picUrl
	r.Set("pic_url", _picUrl)
	return nil
}

// GetPicUrl PicUrl Getter
func (r AlibabaalihealthpregnancypostsdataAPIRequest) GetPicUrl() string {
	return r._picUrl
}

// SetUserId is UserId Setter
// 用户id
func (r *AlibabaalihealthpregnancypostsdataAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaalihealthpregnancypostsdataAPIRequest) GetUserId() int64 {
	return r._userId
}

// SetEventType is EventType Setter
// 事件类型 0发帖 1回帖
func (r *AlibabaalihealthpregnancypostsdataAPIRequest) SetEventType(_eventType int64) error {
	r._eventType = _eventType
	r.Set("event_type", _eventType)
	return nil
}

// GetEventType EventType Getter
func (r AlibabaalihealthpregnancypostsdataAPIRequest) GetEventType() int64 {
	return r._eventType
}

// SetMainId is MainId Setter
// 主贴id
func (r *AlibabaalihealthpregnancypostsdataAPIRequest) SetMainId(_mainId int64) error {
	r._mainId = _mainId
	r.Set("main_id", _mainId)
	return nil
}

// GetMainId MainId Getter
func (r AlibabaalihealthpregnancypostsdataAPIRequest) GetMainId() int64 {
	return r._mainId
}

// SetReplyId is ReplyId Setter
// 回帖id
func (r *AlibabaalihealthpregnancypostsdataAPIRequest) SetReplyId(_replyId int64) error {
	r._replyId = _replyId
	r.Set("reply_id", _replyId)
	return nil
}

// GetReplyId ReplyId Getter
func (r AlibabaalihealthpregnancypostsdataAPIRequest) GetReplyId() int64 {
	return r._replyId
}

// SetDate is Date Setter
// 发帖时间
func (r *AlibabaalihealthpregnancypostsdataAPIRequest) SetDate(_date int64) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r AlibabaalihealthpregnancypostsdataAPIRequest) GetDate() int64 {
	return r._date
}

package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractactivityaddcommentAPIRequest 微淘评论接口 API请求
// alibaba.interact.activity.addcomment
//
// 发表评论，并返回楼层
type AlibabainteractactivityaddcommentAPIRequest struct {
	model.Params
	// 该字段为评论内容
	_content string
	// 发评论的业务id
	_bizId string
	// 评论feedid
	_feedId int64
}

// NewAlibabainteractactivityaddcommentRequest 初始化AlibabainteractactivityaddcommentAPIRequest对象
func NewAlibabainteractactivityaddcommentRequest() *AlibabainteractactivityaddcommentAPIRequest {
	return &AlibabainteractactivityaddcommentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractactivityaddcommentAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.activity.addcomment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractactivityaddcommentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractactivityaddcommentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// 该字段为评论内容
func (r *AlibabainteractactivityaddcommentAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabainteractactivityaddcommentAPIRequest) GetContent() string {
	return r._content
}

// SetBizId is BizId Setter
// 发评论的业务id
func (r *AlibabainteractactivityaddcommentAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// GetBizId BizId Getter
func (r AlibabainteractactivityaddcommentAPIRequest) GetBizId() string {
	return r._bizId
}

// SetFeedId is FeedId Setter
// 评论feedid
func (r *AlibabainteractactivityaddcommentAPIRequest) SetFeedId(_feedId int64) error {
	r._feedId = _feedId
	r.Set("feed_id", _feedId)
	return nil
}

// GetFeedId FeedId Getter
func (r AlibabainteractactivityaddcommentAPIRequest) GetFeedId() int64 {
	return r._feedId
}

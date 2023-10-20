package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractActivityAddcommentAPIRequest 微淘评论接口 API请求
// alibaba.interact.activity.addcomment
//
// 发表评论，并返回楼层
type AlibabaInteractActivityAddcommentAPIRequest struct {
	model.Params
	// 该字段为评论内容
	_content string
	// 发评论的业务id
	_bizId string
	// 评论feedid
	_feedId int64
}

// NewAlibabaInteractActivityAddcommentRequest 初始化AlibabaInteractActivityAddcommentAPIRequest对象
func NewAlibabaInteractActivityAddcommentRequest() *AlibabaInteractActivityAddcommentAPIRequest {
	return &AlibabaInteractActivityAddcommentAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractActivityAddcommentAPIRequest) Reset() {
	r._content = ""
	r._bizId = ""
	r._feedId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractActivityAddcommentAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.activity.addcomment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractActivityAddcommentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractActivityAddcommentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// 该字段为评论内容
func (r *AlibabaInteractActivityAddcommentAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabaInteractActivityAddcommentAPIRequest) GetContent() string {
	return r._content
}

// SetBizId is BizId Setter
// 发评论的业务id
func (r *AlibabaInteractActivityAddcommentAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// GetBizId BizId Getter
func (r AlibabaInteractActivityAddcommentAPIRequest) GetBizId() string {
	return r._bizId
}

// SetFeedId is FeedId Setter
// 评论feedid
func (r *AlibabaInteractActivityAddcommentAPIRequest) SetFeedId(_feedId int64) error {
	r._feedId = _feedId
	r.Set("feed_id", _feedId)
	return nil
}

// GetFeedId FeedId Getter
func (r AlibabaInteractActivityAddcommentAPIRequest) GetFeedId() int64 {
	return r._feedId
}

var poolAlibabaInteractActivityAddcommentAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractActivityAddcommentRequest()
	},
}

// GetAlibabaInteractActivityAddcommentRequest 从 sync.Pool 获取 AlibabaInteractActivityAddcommentAPIRequest
func GetAlibabaInteractActivityAddcommentAPIRequest() *AlibabaInteractActivityAddcommentAPIRequest {
	return poolAlibabaInteractActivityAddcommentAPIRequest.Get().(*AlibabaInteractActivityAddcommentAPIRequest)
}

// ReleaseAlibabaInteractActivityAddcommentAPIRequest 将 AlibabaInteractActivityAddcommentAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractActivityAddcommentAPIRequest(v *AlibabaInteractActivityAddcommentAPIRequest) {
	v.Reset()
	poolAlibabaInteractActivityAddcommentAPIRequest.Put(v)
}

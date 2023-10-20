package dmp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDmpCrowdTemplateTopicFindAPIRequest 平台精选榜单和模版查询接口 API请求
// taobao.dmp.crowd.template.topic.find
//
// 查询平台精选榜单和模版信息
type TaobaoDmpCrowdTemplateTopicFindAPIRequest struct {
	model.Params
	// 请求体
	_apiContext *ApiContextDto
	// 查询对象
	_topicQuery *TopicQueryDto
}

// NewTaobaoDmpCrowdTemplateTopicFindRequest 初始化TaobaoDmpCrowdTemplateTopicFindAPIRequest对象
func NewTaobaoDmpCrowdTemplateTopicFindRequest() *TaobaoDmpCrowdTemplateTopicFindAPIRequest {
	return &TaobaoDmpCrowdTemplateTopicFindAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoDmpCrowdTemplateTopicFindAPIRequest) Reset() {
	r._apiContext = nil
	r._topicQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDmpCrowdTemplateTopicFindAPIRequest) GetApiMethodName() string {
	return "taobao.dmp.crowd.template.topic.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDmpCrowdTemplateTopicFindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDmpCrowdTemplateTopicFindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiContext is ApiContext Setter
// 请求体
func (r *TaobaoDmpCrowdTemplateTopicFindAPIRequest) SetApiContext(_apiContext *ApiContextDto) error {
	r._apiContext = _apiContext
	r.Set("api_context", _apiContext)
	return nil
}

// GetApiContext ApiContext Getter
func (r TaobaoDmpCrowdTemplateTopicFindAPIRequest) GetApiContext() *ApiContextDto {
	return r._apiContext
}

// SetTopicQuery is TopicQuery Setter
// 查询对象
func (r *TaobaoDmpCrowdTemplateTopicFindAPIRequest) SetTopicQuery(_topicQuery *TopicQueryDto) error {
	r._topicQuery = _topicQuery
	r.Set("topic_query", _topicQuery)
	return nil
}

// GetTopicQuery TopicQuery Getter
func (r TaobaoDmpCrowdTemplateTopicFindAPIRequest) GetTopicQuery() *TopicQueryDto {
	return r._topicQuery
}

var poolTaobaoDmpCrowdTemplateTopicFindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoDmpCrowdTemplateTopicFindRequest()
	},
}

// GetTaobaoDmpCrowdTemplateTopicFindRequest 从 sync.Pool 获取 TaobaoDmpCrowdTemplateTopicFindAPIRequest
func GetTaobaoDmpCrowdTemplateTopicFindAPIRequest() *TaobaoDmpCrowdTemplateTopicFindAPIRequest {
	return poolTaobaoDmpCrowdTemplateTopicFindAPIRequest.Get().(*TaobaoDmpCrowdTemplateTopicFindAPIRequest)
}

// ReleaseTaobaoDmpCrowdTemplateTopicFindAPIRequest 将 TaobaoDmpCrowdTemplateTopicFindAPIRequest 放入 sync.Pool
func ReleaseTaobaoDmpCrowdTemplateTopicFindAPIRequest(v *TaobaoDmpCrowdTemplateTopicFindAPIRequest) {
	v.Reset()
	poolTaobaoDmpCrowdTemplateTopicFindAPIRequest.Put(v)
}

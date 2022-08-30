package dmp

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDmpCrowdTemplateTopicFindAPIRequest) GetApiMethodName() string {
	return "taobao.dmp.crowd.template.topic.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDmpCrowdTemplateTopicFindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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

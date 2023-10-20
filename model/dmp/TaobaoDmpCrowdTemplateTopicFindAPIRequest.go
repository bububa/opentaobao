package dmp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodmpcrowdtemplatetopicfindAPIRequest 平台精选榜单和模版查询接口 API请求
// taobao.dmp.crowd.template.topic.find
//
// 查询平台精选榜单和模版信息
type TaobaodmpcrowdtemplatetopicfindAPIRequest struct {
	model.Params
	// 请求体
	_apiContext *ApiContextDto
	// 查询对象
	_topicQuery *TopicQueryDto
}

// NewTaobaodmpcrowdtemplatetopicfindRequest 初始化TaobaodmpcrowdtemplatetopicfindAPIRequest对象
func NewTaobaodmpcrowdtemplatetopicfindRequest() *TaobaodmpcrowdtemplatetopicfindAPIRequest {
	return &TaobaodmpcrowdtemplatetopicfindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodmpcrowdtemplatetopicfindAPIRequest) GetApiMethodName() string {
	return "taobao.dmp.crowd.template.topic.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodmpcrowdtemplatetopicfindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodmpcrowdtemplatetopicfindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiContext is ApiContext Setter
// 请求体
func (r *TaobaodmpcrowdtemplatetopicfindAPIRequest) SetApiContext(_apiContext *ApiContextDto) error {
	r._apiContext = _apiContext
	r.Set("api_context", _apiContext)
	return nil
}

// GetApiContext ApiContext Getter
func (r TaobaodmpcrowdtemplatetopicfindAPIRequest) GetApiContext() *ApiContextDto {
	return r._apiContext
}

// SetTopicQuery is TopicQuery Setter
// 查询对象
func (r *TaobaodmpcrowdtemplatetopicfindAPIRequest) SetTopicQuery(_topicQuery *TopicQueryDto) error {
	r._topicQuery = _topicQuery
	r.Set("topic_query", _topicQuery)
	return nil
}

// GetTopicQuery TopicQuery Getter
func (r TaobaodmpcrowdtemplatetopicfindAPIRequest) GetTopicQuery() *TopicQueryDto {
	return r._topicQuery
}

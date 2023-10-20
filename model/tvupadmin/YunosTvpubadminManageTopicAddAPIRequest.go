package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminmanagetopicaddAPIRequest 新增专题 API请求
// yunos.tvpubadmin.manage.topic.add
//
// 专题新增
type YunostvpubadminmanagetopicaddAPIRequest struct {
	model.Params
	// 新增topic的json信息
	_topicJson string
}

// NewYunostvpubadminmanagetopicaddRequest 初始化YunostvpubadminmanagetopicaddAPIRequest对象
func NewYunostvpubadminmanagetopicaddRequest() *YunostvpubadminmanagetopicaddAPIRequest {
	return &YunostvpubadminmanagetopicaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadminmanagetopicaddAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.topic.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadminmanagetopicaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadminmanagetopicaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopicJson is TopicJson Setter
// 新增topic的json信息
func (r *YunostvpubadminmanagetopicaddAPIRequest) SetTopicJson(_topicJson string) error {
	r._topicJson = _topicJson
	r.Set("topic_json", _topicJson)
	return nil
}

// GetTopicJson TopicJson Getter
func (r YunostvpubadminmanagetopicaddAPIRequest) GetTopicJson() string {
	return r._topicJson
}

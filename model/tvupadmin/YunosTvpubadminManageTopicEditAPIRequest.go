package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminmanagetopiceditAPIRequest 编辑专题 API请求
// yunos.tvpubadmin.manage.topic.edit
//
// 编辑专题
type YunostvpubadminmanagetopiceditAPIRequest struct {
	model.Params
	// 待编辑专题
	_topicJson string
}

// NewYunostvpubadminmanagetopiceditRequest 初始化YunostvpubadminmanagetopiceditAPIRequest对象
func NewYunostvpubadminmanagetopiceditRequest() *YunostvpubadminmanagetopiceditAPIRequest {
	return &YunostvpubadminmanagetopiceditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadminmanagetopiceditAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.topic.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadminmanagetopiceditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadminmanagetopiceditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopicJson is TopicJson Setter
// 待编辑专题
func (r *YunostvpubadminmanagetopiceditAPIRequest) SetTopicJson(_topicJson string) error {
	r._topicJson = _topicJson
	r.Set("topic_json", _topicJson)
	return nil
}

// GetTopicJson TopicJson Getter
func (r YunostvpubadminmanagetopiceditAPIRequest) GetTopicJson() string {
	return r._topicJson
}

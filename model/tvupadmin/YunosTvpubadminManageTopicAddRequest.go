package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增专题 API请求
yunos.tvpubadmin.manage.topic.add

专题新增
*/
type YunosTvpubadminManageTopicAddRequest struct {
    model.Params
    // 新增topic的json信息
    _topicJson   string
}

// 初始化YunosTvpubadminManageTopicAddRequest对象
func NewYunosTvpubadminManageTopicAddRequest() *YunosTvpubadminManageTopicAddRequest{
    return &YunosTvpubadminManageTopicAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicAddRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.topic.add"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopicJson Setter
// 新增topic的json信息
func (r *YunosTvpubadminManageTopicAddRequest) SetTopicJson(_topicJson string) error {
    r._topicJson = _topicJson
    r.Set("topic_json", _topicJson)
    return nil
}

// TopicJson Getter
func (r YunosTvpubadminManageTopicAddRequest) GetTopicJson() string {
    return r._topicJson
}

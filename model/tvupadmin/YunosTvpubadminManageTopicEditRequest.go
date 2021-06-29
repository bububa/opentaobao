package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑专题 API请求
yunos.tvpubadmin.manage.topic.edit

编辑专题
*/
type YunosTvpubadminManageTopicEditRequest struct {
    model.Params
    // 待编辑专题
    _topicJson   string
}

// 初始化YunosTvpubadminManageTopicEditRequest对象
func NewYunosTvpubadminManageTopicEditRequest() *YunosTvpubadminManageTopicEditRequest{
    return &YunosTvpubadminManageTopicEditRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicEditRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.topic.edit"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopicJson Setter
// 待编辑专题
func (r *YunosTvpubadminManageTopicEditRequest) SetTopicJson(_topicJson string) error {
    r._topicJson = _topicJson
    r.Set("topic_json", _topicJson)
    return nil
}

// TopicJson Getter
func (r YunosTvpubadminManageTopicEditRequest) GetTopicJson() string {
    return r._topicJson
}

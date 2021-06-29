package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增专题 APIRequest
yunos.tvpubadmin.manage.topic.add

专题新增
*/
type YunosTvpubadminManageTopicAddRequest struct {
    model.Params

    // 新增topic的json信息
    topicJson   string 

}

func NewYunosTvpubadminManageTopicAddRequest() *YunosTvpubadminManageTopicAddRequest{
    return &YunosTvpubadminManageTopicAddRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminManageTopicAddRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.topic.add"
}

func (r YunosTvpubadminManageTopicAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminManageTopicAddRequest) SetTopicJson(topicJson string) error {
    r.topicJson = topicJson
    r.Set("topic_json", topicJson)
    return nil
}

func (r YunosTvpubadminManageTopicAddRequest) GetTopicJson() string {
    return r.topicJson
}


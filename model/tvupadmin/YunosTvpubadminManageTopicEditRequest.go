package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑专题 APIRequest
yunos.tvpubadmin.manage.topic.edit

编辑专题
*/
type YunosTvpubadminManageTopicEditRequest struct {
    model.Params

    // 待编辑专题
    topicJson   string 

}

func NewYunosTvpubadminManageTopicEditRequest() *YunosTvpubadminManageTopicEditRequest{
    return &YunosTvpubadminManageTopicEditRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminManageTopicEditRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.topic.edit"
}

func (r YunosTvpubadminManageTopicEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminManageTopicEditRequest) SetTopicJson(topicJson string) error {
    r.topicJson = topicJson
    r.Set("topic_json", topicJson)
    return nil
}

func (r YunosTvpubadminManageTopicEditRequest) GetTopicJson() string {
    return r.topicJson
}


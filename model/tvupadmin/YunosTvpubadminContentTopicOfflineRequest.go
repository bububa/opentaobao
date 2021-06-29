package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松专题下线 APIRequest
yunos.tvpubadmin.content.topic.offline

迎客松专题下线
*/
type YunosTvpubadminContentTopicOfflineRequest struct {
    model.Params

    // id
    id   int64 

}

func NewYunosTvpubadminContentTopicOfflineRequest() *YunosTvpubadminContentTopicOfflineRequest{
    return &YunosTvpubadminContentTopicOfflineRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentTopicOfflineRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.topic.offline"
}

func (r YunosTvpubadminContentTopicOfflineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentTopicOfflineRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r YunosTvpubadminContentTopicOfflineRequest) GetId() int64 {
    return r.id
}


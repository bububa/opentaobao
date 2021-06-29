package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松影视频道下线 APIRequest
yunos.tvpubadmin.content.channel.offline

迎客松影视频道下线
*/
type YunosTvpubadminContentChannelOfflineRequest struct {
    model.Params

    // id
    id   int64 

}

func NewYunosTvpubadminContentChannelOfflineRequest() *YunosTvpubadminContentChannelOfflineRequest{
    return &YunosTvpubadminContentChannelOfflineRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentChannelOfflineRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.channel.offline"
}

func (r YunosTvpubadminContentChannelOfflineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentChannelOfflineRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r YunosTvpubadminContentChannelOfflineRequest) GetId() int64 {
    return r.id
}


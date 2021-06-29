package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松影视频道下线 API请求
yunos.tvpubadmin.content.channel.offline

迎客松影视频道下线
*/
type YunosTvpubadminContentChannelOfflineRequest struct {
    model.Params
    // id
    _id   int64
}

// 初始化YunosTvpubadminContentChannelOfflineRequest对象
func NewYunosTvpubadminContentChannelOfflineRequest() *YunosTvpubadminContentChannelOfflineRequest{
    return &YunosTvpubadminContentChannelOfflineRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChannelOfflineRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.channel.offline"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChannelOfflineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// id
func (r *YunosTvpubadminContentChannelOfflineRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r YunosTvpubadminContentChannelOfflineRequest) GetId() int64 {
    return r._id
}

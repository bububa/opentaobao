package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松专题下线 API请求
yunos.tvpubadmin.content.topic.offline

迎客松专题下线
*/
type YunosTvpubadminContentTopicOfflineRequest struct {
    model.Params
    // id
    _id   int64
}

// 初始化YunosTvpubadminContentTopicOfflineRequest对象
func NewYunosTvpubadminContentTopicOfflineRequest() *YunosTvpubadminContentTopicOfflineRequest{
    return &YunosTvpubadminContentTopicOfflineRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentTopicOfflineRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.topic.offline"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentTopicOfflineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// id
func (r *YunosTvpubadminContentTopicOfflineRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r YunosTvpubadminContentTopicOfflineRequest) GetId() int64 {
    return r._id
}

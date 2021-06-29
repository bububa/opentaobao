package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据id获取专题信息 API请求
yunos.tvpubadmin.manage.topic.findbyid

根据id获取专题信息
*/
type YunosTvpubadminManageTopicFindbyidRequest struct {
    model.Params
    // 专题id
    _id   int64
}

// 初始化YunosTvpubadminManageTopicFindbyidRequest对象
func NewYunosTvpubadminManageTopicFindbyidRequest() *YunosTvpubadminManageTopicFindbyidRequest{
    return &YunosTvpubadminManageTopicFindbyidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicFindbyidRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.topic.findbyid"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicFindbyidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 专题id
func (r *YunosTvpubadminManageTopicFindbyidRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r YunosTvpubadminManageTopicFindbyidRequest) GetId() int64 {
    return r._id
}

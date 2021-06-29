package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据id获取专题信息 APIRequest
yunos.tvpubadmin.manage.topic.findbyid

根据id获取专题信息
*/
type YunosTvpubadminManageTopicFindbyidRequest struct {
    model.Params

    // 专题id
    id   int64 

}

func NewYunosTvpubadminManageTopicFindbyidRequest() *YunosTvpubadminManageTopicFindbyidRequest{
    return &YunosTvpubadminManageTopicFindbyidRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminManageTopicFindbyidRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.topic.findbyid"
}

func (r YunosTvpubadminManageTopicFindbyidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminManageTopicFindbyidRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r YunosTvpubadminManageTopicFindbyidRequest) GetId() int64 {
    return r.id
}


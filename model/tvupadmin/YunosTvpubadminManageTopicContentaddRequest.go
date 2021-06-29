package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
专题新增内容 APIRequest
yunos.tvpubadmin.manage.topic.contentadd

专题新增内容
*/
type YunosTvpubadminManageTopicContentaddRequest struct {
    model.Params

    // 新增的专题内容
    contentJson   string 

}

func NewYunosTvpubadminManageTopicContentaddRequest() *YunosTvpubadminManageTopicContentaddRequest{
    return &YunosTvpubadminManageTopicContentaddRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminManageTopicContentaddRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.topic.contentadd"
}

func (r YunosTvpubadminManageTopicContentaddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminManageTopicContentaddRequest) SetContentJson(contentJson string) error {
    r.contentJson = contentJson
    r.Set("content_json", contentJson)
    return nil
}

func (r YunosTvpubadminManageTopicContentaddRequest) GetContentJson() string {
    return r.contentJson
}


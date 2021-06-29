package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
专题关联内容编辑 APIRequest
yunos.tvpubadmin.manage.topic.contentedit

编辑专题关联的内容
*/
type YunosTvpubadminManageTopicContenteditRequest struct {
    model.Params

    // 入参，json格式
    contentJson   string 

}

func NewYunosTvpubadminManageTopicContenteditRequest() *YunosTvpubadminManageTopicContenteditRequest{
    return &YunosTvpubadminManageTopicContenteditRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminManageTopicContenteditRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.topic.contentedit"
}

func (r YunosTvpubadminManageTopicContenteditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminManageTopicContenteditRequest) SetContentJson(contentJson string) error {
    r.contentJson = contentJson
    r.Set("content_json", contentJson)
    return nil
}

func (r YunosTvpubadminManageTopicContenteditRequest) GetContentJson() string {
    return r.contentJson
}


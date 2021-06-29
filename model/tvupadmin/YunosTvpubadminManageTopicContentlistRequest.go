package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查看专题内容列表 APIRequest
yunos.tvpubadmin.manage.topic.contentlist

查看专题内容列表
*/
type YunosTvpubadminManageTopicContentlistRequest struct {
    model.Params

    // 节目查询参数
    programQuery   string 

}

func NewYunosTvpubadminManageTopicContentlistRequest() *YunosTvpubadminManageTopicContentlistRequest{
    return &YunosTvpubadminManageTopicContentlistRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminManageTopicContentlistRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.topic.contentlist"
}

func (r YunosTvpubadminManageTopicContentlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminManageTopicContentlistRequest) SetProgramQuery(programQuery string) error {
    r.programQuery = programQuery
    r.Set("program_query", programQuery)
    return nil
}

func (r YunosTvpubadminManageTopicContentlistRequest) GetProgramQuery() string {
    return r.programQuery
}


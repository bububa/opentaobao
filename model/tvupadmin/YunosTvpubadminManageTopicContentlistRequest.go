package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查看专题内容列表 API请求
yunos.tvpubadmin.manage.topic.contentlist

查看专题内容列表
*/
type YunosTvpubadminManageTopicContentlistRequest struct {
    model.Params
    // 节目查询参数
    _programQuery   string
}

// 初始化YunosTvpubadminManageTopicContentlistRequest对象
func NewYunosTvpubadminManageTopicContentlistRequest() *YunosTvpubadminManageTopicContentlistRequest{
    return &YunosTvpubadminManageTopicContentlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicContentlistRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.topic.contentlist"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicContentlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProgramQuery Setter
// 节目查询参数
func (r *YunosTvpubadminManageTopicContentlistRequest) SetProgramQuery(_programQuery string) error {
    r._programQuery = _programQuery
    r.Set("program_query", _programQuery)
    return nil
}

// ProgramQuery Getter
func (r YunosTvpubadminManageTopicContentlistRequest) GetProgramQuery() string {
    return r._programQuery
}

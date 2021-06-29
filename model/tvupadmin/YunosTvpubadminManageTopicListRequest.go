package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
专题内容操作列表 API请求
yunos.tvpubadmin.manage.topic.list

获取外部可操作编辑的专题列表
*/
type YunosTvpubadminManageTopicListRequest struct {
    model.Params
    // 查询条件
    query   string
}

// 初始化YunosTvpubadminManageTopicListRequest对象
func NewYunosTvpubadminManageTopicListRequest() *YunosTvpubadminManageTopicListRequest{
    return &YunosTvpubadminManageTopicListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicListRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.topic.list"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 查询条件
func (r *YunosTvpubadminManageTopicListRequest) SetQuery(query string) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r YunosTvpubadminManageTopicListRequest) GetQuery() string {
    return r.query
}

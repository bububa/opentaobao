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
type YunosTvpubadminManageTopicListAPIRequest struct {
    model.Params
    // 查询条件
    _query   string
}

// 初始化YunosTvpubadminManageTopicListAPIRequest对象
func NewYunosTvpubadminManageTopicListRequest() *YunosTvpubadminManageTopicListAPIRequest{
    return &YunosTvpubadminManageTopicListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicListAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.topic.list"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 查询条件
func (r *YunosTvpubadminManageTopicListAPIRequest) SetQuery(_query string) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r YunosTvpubadminManageTopicListAPIRequest) GetQuery() string {
    return r._query
}

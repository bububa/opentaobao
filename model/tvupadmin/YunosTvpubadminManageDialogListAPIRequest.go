package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取弹窗列表 API请求
yunos.tvpubadmin.manage.dialog.list

分页获取弹窗配置列表
*/
type YunosTvpubadminManageDialogListAPIRequest struct {
    model.Params
    // 查询的query参数
    _query   string
}

// 初始化YunosTvpubadminManageDialogListAPIRequest对象
func NewYunosTvpubadminManageDialogListRequest() *YunosTvpubadminManageDialogListAPIRequest{
    return &YunosTvpubadminManageDialogListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageDialogListAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.dialog.list"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageDialogListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 查询的query参数
func (r *YunosTvpubadminManageDialogListAPIRequest) SetQuery(_query string) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r YunosTvpubadminManageDialogListAPIRequest) GetQuery() string {
    return r._query
}

package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松查看小酷宝桌面坑位元数据列表 API请求
yunos.tvpubadmin.content.tableaudit.querychilddesktop

迎客松查看小酷宝桌面坑位元数据列表
*/
type YunosTvpubadminContentTableauditQuerychilddesktopRequest struct {
    model.Params
    // 小酷宝桌面坑位查询参数
    _query   string
}

// 初始化YunosTvpubadminContentTableauditQuerychilddesktopRequest对象
func NewYunosTvpubadminContentTableauditQuerychilddesktopRequest() *YunosTvpubadminContentTableauditQuerychilddesktopRequest{
    return &YunosTvpubadminContentTableauditQuerychilddesktopRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentTableauditQuerychilddesktopRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.tableaudit.querychilddesktop"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentTableauditQuerychilddesktopRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 小酷宝桌面坑位查询参数
func (r *YunosTvpubadminContentTableauditQuerychilddesktopRequest) SetQuery(_query string) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r YunosTvpubadminContentTableauditQuerychilddesktopRequest) GetQuery() string {
    return r._query
}

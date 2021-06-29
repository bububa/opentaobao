package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按牌照查询应用 API请求
yunos.tvpubadmin.content.app.querybylicence

按牌照查询应用
*/
type YunosTvpubadminContentAppQuerybylicenceRequest struct {
    model.Params
    // 查询条件
    _query   string
}

// 初始化YunosTvpubadminContentAppQuerybylicenceRequest对象
func NewYunosTvpubadminContentAppQuerybylicenceRequest() *YunosTvpubadminContentAppQuerybylicenceRequest{
    return &YunosTvpubadminContentAppQuerybylicenceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentAppQuerybylicenceRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.app.querybylicence"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentAppQuerybylicenceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 查询条件
func (r *YunosTvpubadminContentAppQuerybylicenceRequest) SetQuery(_query string) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r YunosTvpubadminContentAppQuerybylicenceRequest) GetQuery() string {
    return r._query
}

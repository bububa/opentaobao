package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优酷OTT端广告素材查询 API请求
yunos.tvpubadmin.adm.ott.query

查询广告素材
*/
type YunosTvpubadminAdmOttQueryRequest struct {
    model.Params
    // 查询参数json格式
    _query   string
}

// 初始化YunosTvpubadminAdmOttQueryRequest对象
func NewYunosTvpubadminAdmOttQueryRequest() *YunosTvpubadminAdmOttQueryRequest{
    return &YunosTvpubadminAdmOttQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminAdmOttQueryRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.adm.ott.query"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminAdmOttQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 查询参数json格式
func (r *YunosTvpubadminAdmOttQueryRequest) SetQuery(_query string) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r YunosTvpubadminAdmOttQueryRequest) GetQuery() string {
    return r._query
}

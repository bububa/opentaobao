package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询应用信息 APIRequest
yunos.tvpubadmin.content.app.queryapp

查询应用信息
*/
type YunosTvpubadminContentAppQueryappRequest struct {
    model.Params

    // 查询条件
    query   string 

}

func NewYunosTvpubadminContentAppQueryappRequest() *YunosTvpubadminContentAppQueryappRequest{
    return &YunosTvpubadminContentAppQueryappRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentAppQueryappRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.app.queryapp"
}

func (r YunosTvpubadminContentAppQueryappRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentAppQueryappRequest) SetQuery(query string) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r YunosTvpubadminContentAppQueryappRequest) GetQuery() string {
    return r.query
}


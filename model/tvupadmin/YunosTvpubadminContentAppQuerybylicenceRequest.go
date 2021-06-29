package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按牌照查询应用 APIRequest
yunos.tvpubadmin.content.app.querybylicence

按牌照查询应用
*/
type YunosTvpubadminContentAppQuerybylicenceRequest struct {
    model.Params

    // 查询条件
    query   string 

}

func NewYunosTvpubadminContentAppQuerybylicenceRequest() *YunosTvpubadminContentAppQuerybylicenceRequest{
    return &YunosTvpubadminContentAppQuerybylicenceRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentAppQuerybylicenceRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.app.querybylicence"
}

func (r YunosTvpubadminContentAppQuerybylicenceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentAppQuerybylicenceRequest) SetQuery(query string) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r YunosTvpubadminContentAppQuerybylicenceRequest) GetQuery() string {
    return r.query
}


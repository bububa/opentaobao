package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优酷OTT端广告素材查询 APIRequest
yunos.tvpubadmin.adm.ott.query

查询广告素材
*/
type YunosTvpubadminAdmOttQueryRequest struct {
    model.Params

    // 查询参数json格式
    query   string 

}

func NewYunosTvpubadminAdmOttQueryRequest() *YunosTvpubadminAdmOttQueryRequest{
    return &YunosTvpubadminAdmOttQueryRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminAdmOttQueryRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.adm.ott.query"
}

func (r YunosTvpubadminAdmOttQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminAdmOttQueryRequest) SetQuery(query string) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r YunosTvpubadminAdmOttQueryRequest) GetQuery() string {
    return r.query
}


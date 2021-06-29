package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松视频审核记录查询 APIRequest
yunos.tvpubadmin.content.video.getauditlist

迎客松视频审核记录查询
*/
type YunosTvpubadminContentVideoGetauditlistRequest struct {
    model.Params

    // 查询
    query   string 

}

func NewYunosTvpubadminContentVideoGetauditlistRequest() *YunosTvpubadminContentVideoGetauditlistRequest{
    return &YunosTvpubadminContentVideoGetauditlistRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentVideoGetauditlistRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.video.getauditlist"
}

func (r YunosTvpubadminContentVideoGetauditlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentVideoGetauditlistRequest) SetQuery(query string) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r YunosTvpubadminContentVideoGetauditlistRequest) GetQuery() string {
    return r.query
}


package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
广告牌照管控修改 APIRequest
yunos.tvpubadmin.content.advert.manageschedule

广告牌照管控修改
*/
type YunosTvpubadminContentAdvertManagescheduleRequest struct {
    model.Params

    // 管理参数
    req   string 

}

func NewYunosTvpubadminContentAdvertManagescheduleRequest() *YunosTvpubadminContentAdvertManagescheduleRequest{
    return &YunosTvpubadminContentAdvertManagescheduleRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentAdvertManagescheduleRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.advert.manageschedule"
}

func (r YunosTvpubadminContentAdvertManagescheduleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentAdvertManagescheduleRequest) SetReq(req string) error {
    r.req = req
    r.Set("req", req)
    return nil
}

func (r YunosTvpubadminContentAdvertManagescheduleRequest) GetReq() string {
    return r.req
}


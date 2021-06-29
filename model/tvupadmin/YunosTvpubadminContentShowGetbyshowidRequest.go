package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松根据节目id获取节目元数据 APIRequest
yunos.tvpubadmin.content.show.getbyshowid

迎客松根据节目id获取节目元数据
*/
type YunosTvpubadminContentShowGetbyshowidRequest struct {
    model.Params

    // 节目字符串id
    showId   string 

}

func NewYunosTvpubadminContentShowGetbyshowidRequest() *YunosTvpubadminContentShowGetbyshowidRequest{
    return &YunosTvpubadminContentShowGetbyshowidRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentShowGetbyshowidRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.show.getbyshowid"
}

func (r YunosTvpubadminContentShowGetbyshowidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentShowGetbyshowidRequest) SetShowId(showId string) error {
    r.showId = showId
    r.Set("show_id", showId)
    return nil
}

func (r YunosTvpubadminContentShowGetbyshowidRequest) GetShowId() string {
    return r.showId
}


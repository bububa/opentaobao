package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松根据节目longid获取节目元数据 APIRequest
yunos.tvpubadmin.content.show.getbyshowlongid

迎客松根据节目longid获取节目元数据
*/
type YunosTvpubadminContentShowGetbyshowlongidRequest struct {
    model.Params

    // 节目longid
    showLongId   int64 

}

func NewYunosTvpubadminContentShowGetbyshowlongidRequest() *YunosTvpubadminContentShowGetbyshowlongidRequest{
    return &YunosTvpubadminContentShowGetbyshowlongidRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentShowGetbyshowlongidRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.show.getbyshowlongid"
}

func (r YunosTvpubadminContentShowGetbyshowlongidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentShowGetbyshowlongidRequest) SetShowLongId(showLongId int64) error {
    r.showLongId = showLongId
    r.Set("show_long_id", showLongId)
    return nil
}

func (r YunosTvpubadminContentShowGetbyshowlongidRequest) GetShowLongId() int64 {
    return r.showLongId
}


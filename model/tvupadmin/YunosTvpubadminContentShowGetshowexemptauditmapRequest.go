package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松批量查询节目某个牌照的免审状态 APIRequest
yunos.tvpubadmin.content.show.getshowexemptauditmap

迎客松批量查询节目某个牌照的免审状态
*/
type YunosTvpubadminContentShowGetshowexemptauditmapRequest struct {
    model.Params

    // 节目longid
    showLongIds   string 

    // 牌照id
    license   int64 

}

func NewYunosTvpubadminContentShowGetshowexemptauditmapRequest() *YunosTvpubadminContentShowGetshowexemptauditmapRequest{
    return &YunosTvpubadminContentShowGetshowexemptauditmapRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentShowGetshowexemptauditmapRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.show.getshowexemptauditmap"
}

func (r YunosTvpubadminContentShowGetshowexemptauditmapRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentShowGetshowexemptauditmapRequest) SetShowLongIds(showLongIds string) error {
    r.showLongIds = showLongIds
    r.Set("show_long_ids", showLongIds)
    return nil
}

func (r YunosTvpubadminContentShowGetshowexemptauditmapRequest) GetShowLongIds() string {
    return r.showLongIds
}

func (r *YunosTvpubadminContentShowGetshowexemptauditmapRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminContentShowGetshowexemptauditmapRequest) GetLicense() int64 {
    return r.license
}


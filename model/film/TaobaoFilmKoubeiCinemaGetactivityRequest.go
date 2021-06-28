package film

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
口碑-影院营销数据查询 APIRequest
taobao.film.koubei.cinema.getactivity

口碑-影院营销数据查询
*/
type TaobaoFilmKoubeiCinemaGetactivityRequest struct {
    model.Params

    // 用户账号
    userId   string 

    // 账号类型
    accountType   string 

    // 城市编码
    cityCode   int64 

    // 平台
    platform   int64 

    // 影院ID集合
    cinemaIds   []int64 

    // 附加参数
    params   string 

}

func NewTaobaoFilmKoubeiCinemaGetactivityRequest() *TaobaoFilmKoubeiCinemaGetactivityRequest{
    return &TaobaoFilmKoubeiCinemaGetactivityRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFilmKoubeiCinemaGetactivityRequest) GetApiMethodName() string {
    return "taobao.film.koubei.cinema.getactivity"
}

func (r TaobaoFilmKoubeiCinemaGetactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFilmKoubeiCinemaGetactivityRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoFilmKoubeiCinemaGetactivityRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoFilmKoubeiCinemaGetactivityRequest) SetAccountType(accountType string) error {
    r.accountType = accountType
    r.Set("account_type", accountType)
    return nil
}

func (r TaobaoFilmKoubeiCinemaGetactivityRequest) GetAccountType() string {
    return r.accountType
}

func (r *TaobaoFilmKoubeiCinemaGetactivityRequest) SetCityCode(cityCode int64) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r TaobaoFilmKoubeiCinemaGetactivityRequest) GetCityCode() int64 {
    return r.cityCode
}

func (r *TaobaoFilmKoubeiCinemaGetactivityRequest) SetPlatform(platform int64) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

func (r TaobaoFilmKoubeiCinemaGetactivityRequest) GetPlatform() int64 {
    return r.platform
}

func (r *TaobaoFilmKoubeiCinemaGetactivityRequest) SetCinemaIds(cinemaIds []int64) error {
    r.cinemaIds = cinemaIds
    r.Set("cinema_ids", cinemaIds)
    return nil
}

func (r TaobaoFilmKoubeiCinemaGetactivityRequest) GetCinemaIds() []int64 {
    return r.cinemaIds
}

func (r *TaobaoFilmKoubeiCinemaGetactivityRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

func (r TaobaoFilmKoubeiCinemaGetactivityRequest) GetParams() string {
    return r.params
}


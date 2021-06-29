package film

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
口碑-影院营销数据查询 API请求
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

// 初始化TaobaoFilmKoubeiCinemaGetactivityRequest对象
func NewTaobaoFilmKoubeiCinemaGetactivityRequest() *TaobaoFilmKoubeiCinemaGetactivityRequest{
    return &TaobaoFilmKoubeiCinemaGetactivityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFilmKoubeiCinemaGetactivityRequest) GetApiMethodName() string {
    return "taobao.film.koubei.cinema.getactivity"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFilmKoubeiCinemaGetactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户账号
func (r *TaobaoFilmKoubeiCinemaGetactivityRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoFilmKoubeiCinemaGetactivityRequest) GetUserId() string {
    return r.userId
}
// AccountType Setter
// 账号类型
func (r *TaobaoFilmKoubeiCinemaGetactivityRequest) SetAccountType(accountType string) error {
    r.accountType = accountType
    r.Set("account_type", accountType)
    return nil
}

// AccountType Getter
func (r TaobaoFilmKoubeiCinemaGetactivityRequest) GetAccountType() string {
    return r.accountType
}
// CityCode Setter
// 城市编码
func (r *TaobaoFilmKoubeiCinemaGetactivityRequest) SetCityCode(cityCode int64) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

// CityCode Getter
func (r TaobaoFilmKoubeiCinemaGetactivityRequest) GetCityCode() int64 {
    return r.cityCode
}
// Platform Setter
// 平台
func (r *TaobaoFilmKoubeiCinemaGetactivityRequest) SetPlatform(platform int64) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

// Platform Getter
func (r TaobaoFilmKoubeiCinemaGetactivityRequest) GetPlatform() int64 {
    return r.platform
}
// CinemaIds Setter
// 影院ID集合
func (r *TaobaoFilmKoubeiCinemaGetactivityRequest) SetCinemaIds(cinemaIds []int64) error {
    r.cinemaIds = cinemaIds
    r.Set("cinema_ids", cinemaIds)
    return nil
}

// CinemaIds Getter
func (r TaobaoFilmKoubeiCinemaGetactivityRequest) GetCinemaIds() []int64 {
    return r.cinemaIds
}
// Params Setter
// 附加参数
func (r *TaobaoFilmKoubeiCinemaGetactivityRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

// Params Getter
func (r TaobaoFilmKoubeiCinemaGetactivityRequest) GetParams() string {
    return r.params
}

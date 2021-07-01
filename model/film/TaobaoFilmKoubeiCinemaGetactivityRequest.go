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
type TaobaoFilmKoubeiCinemaGetactivityAPIRequest struct {
    model.Params
    // 用户账号
    _userId   string
    // 账号类型
    _accountType   string
    // 城市编码
    _cityCode   int64
    // 平台
    _platform   int64
    // 影院ID集合
    _cinemaIds   []int64
    // 附加参数
    _params   string
}

// 初始化TaobaoFilmKoubeiCinemaGetactivityAPIRequest对象
func NewTaobaoFilmKoubeiCinemaGetactivityRequest() *TaobaoFilmKoubeiCinemaGetactivityAPIRequest{
    return &TaobaoFilmKoubeiCinemaGetactivityAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFilmKoubeiCinemaGetactivityAPIRequest) GetApiMethodName() string {
    return "taobao.film.koubei.cinema.getactivity"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFilmKoubeiCinemaGetactivityAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户账号
func (r *TaobaoFilmKoubeiCinemaGetactivityAPIRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoFilmKoubeiCinemaGetactivityAPIRequest) GetUserId() string {
    return r._userId
}
// AccountType Setter
// 账号类型
func (r *TaobaoFilmKoubeiCinemaGetactivityAPIRequest) SetAccountType(_accountType string) error {
    r._accountType = _accountType
    r.Set("account_type", _accountType)
    return nil
}

// AccountType Getter
func (r TaobaoFilmKoubeiCinemaGetactivityAPIRequest) GetAccountType() string {
    return r._accountType
}
// CityCode Setter
// 城市编码
func (r *TaobaoFilmKoubeiCinemaGetactivityAPIRequest) SetCityCode(_cityCode int64) error {
    r._cityCode = _cityCode
    r.Set("city_code", _cityCode)
    return nil
}

// CityCode Getter
func (r TaobaoFilmKoubeiCinemaGetactivityAPIRequest) GetCityCode() int64 {
    return r._cityCode
}
// Platform Setter
// 平台
func (r *TaobaoFilmKoubeiCinemaGetactivityAPIRequest) SetPlatform(_platform int64) error {
    r._platform = _platform
    r.Set("platform", _platform)
    return nil
}

// Platform Getter
func (r TaobaoFilmKoubeiCinemaGetactivityAPIRequest) GetPlatform() int64 {
    return r._platform
}
// CinemaIds Setter
// 影院ID集合
func (r *TaobaoFilmKoubeiCinemaGetactivityAPIRequest) SetCinemaIds(_cinemaIds []int64) error {
    r._cinemaIds = _cinemaIds
    r.Set("cinema_ids", _cinemaIds)
    return nil
}

// CinemaIds Getter
func (r TaobaoFilmKoubeiCinemaGetactivityAPIRequest) GetCinemaIds() []int64 {
    return r._cinemaIds
}
// Params Setter
// 附加参数
func (r *TaobaoFilmKoubeiCinemaGetactivityAPIRequest) SetParams(_params string) error {
    r._params = _params
    r.Set("params", _params)
    return nil
}

// Params Getter
func (r TaobaoFilmKoubeiCinemaGetactivityAPIRequest) GetParams() string {
    return r._params
}

package film

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/film"
)

/* 
口碑-影院营销数据查询 
taobao.film.koubei.cinema.getactivity

口碑-影院营销数据查询
*/
func TaobaoFilmKoubeiCinemaGetactivity(clt *core.SDKClient, req *film.TaobaoFilmKoubeiCinemaGetactivityRequest, session string) (*film.TaobaoFilmKoubeiCinemaGetactivityResponse, error) {
    var resp film.TaobaoFilmKoubeiCinemaGetactivityAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

package film

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/film"
)

/* 
淘票票外部直发券 
taobao.film.lottery.sendcode

淘票票外部直发券
*/
func TaobaoFilmLotterySendcode(clt *core.SDKClient, req *film.TaobaoFilmLotterySendcodeRequest, session string) (*film.TaobaoFilmLotterySendcodeResponse, error) {
    var resp film.TaobaoFilmLotterySendcodeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

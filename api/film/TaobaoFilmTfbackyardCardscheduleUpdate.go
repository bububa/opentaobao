package film

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/film"
)

/* 
CGV影城卡排期数据传输 
taobao.film.tfbackyard.cardschedule.update

cgv影城卡排期价格数据传输API
*/
func TaobaoFilmTfbackyardCardscheduleUpdate(clt *core.SDKClient, req *film.TaobaoFilmTfbackyardCardscheduleUpdateRequest, session string) (*film.TaobaoFilmTfbackyardCardscheduleUpdateAPIResponse, error) {
    var resp film.TaobaoFilmTfbackyardCardscheduleUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

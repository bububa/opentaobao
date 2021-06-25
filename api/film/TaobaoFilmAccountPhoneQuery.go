package film

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/film"
)

/* 
根据手机查询匹配账号列表 
taobao.film.account.phone.query

根据手机号查询匹配的账号列表
*/
func TaobaoFilmAccountPhoneQuery(clt *core.SDKClient, req *film.TaobaoFilmAccountPhoneQueryRequest, session string) (*film.TaobaoFilmAccountPhoneQueryResponse, error) {
    var resp film.TaobaoFilmAccountPhoneQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

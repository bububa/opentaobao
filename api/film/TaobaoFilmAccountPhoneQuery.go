package film

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/film"
)

// TaobaoFilmAccountPhoneQuery 根据手机查询匹配账号列表
// taobao.film.account.phone.query
//
// 根据手机号查询匹配的账号列表
func TaobaoFilmAccountPhoneQuery(clt *core.SDKClient, req *film.TaobaoFilmAccountPhoneQueryAPIRequest, session string) (*film.TaobaoFilmAccountPhoneQueryAPIResponse, error) {
	var resp film.TaobaoFilmAccountPhoneQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

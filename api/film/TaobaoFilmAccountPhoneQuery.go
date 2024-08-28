package film

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/film"
)

// TaobaoFilmAccountPhoneQuery 根据手机查询匹配账号列表
// taobao.film.account.phone.query
//
// 根据手机号查询匹配的账号列表
func TaobaoFilmAccountPhoneQuery(ctx context.Context, clt *core.SDKClient, req *film.TaobaoFilmAccountPhoneQueryAPIRequest, resp *film.TaobaoFilmAccountPhoneQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

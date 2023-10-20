package film

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/film"
)

// Taobaofilmaccountphonequery 根据手机查询匹配账号列表
// taobao.film.account.phone.query
//
// 根据手机号查询匹配的账号列表
func Taobaofilmaccountphonequery(clt *core.SDKClient, req *film.TaobaofilmaccountphonequeryAPIRequest, session string) (*film.TaobaofilmaccountphonequeryAPIResponse, error) {
	var resp film.TaobaofilmaccountphonequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

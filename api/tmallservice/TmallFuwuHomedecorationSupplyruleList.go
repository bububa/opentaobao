package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallfuwuhomedecorationsupplyrulelist 查询供给规则接口
// tmall.fuwu.homedecoration.supplyrule.list
//
// 查询供给规则接口
func Tmallfuwuhomedecorationsupplyrulelist(clt *core.SDKClient, req *tmallservice.TmallfuwuhomedecorationsupplyrulelistAPIRequest, session string) (*tmallservice.TmallfuwuhomedecorationsupplyrulelistAPIResponse, error) {
	var resp tmallservice.TmallfuwuhomedecorationsupplyrulelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

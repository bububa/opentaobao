package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallfuwuhomedecorationsupplyrulecategoryworkerlist 基于规则查品牌品类工人接口
// tmall.fuwu.homedecoration.supplyrule.categoryworkerlist
//
// 基于规则查品牌品类工人接口
func Tmallfuwuhomedecorationsupplyrulecategoryworkerlist(clt *core.SDKClient, req *tmallservice.TmallfuwuhomedecorationsupplyrulecategoryworkerlistAPIRequest, session string) (*tmallservice.TmallfuwuhomedecorationsupplyrulecategoryworkerlistAPIResponse, error) {
	var resp tmallservice.TmallfuwuhomedecorationsupplyrulecategoryworkerlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

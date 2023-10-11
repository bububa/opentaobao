package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallFuwuHomedecorationSupplyruleCategoryworkerlist 基于规则查品牌品类工人接口
// tmall.fuwu.homedecoration.supplyrule.categoryworkerlist
//
// 基于规则查品牌品类工人接口
func TmallFuwuHomedecorationSupplyruleCategoryworkerlist(clt *core.SDKClient, req *tmallservice.TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest, session string) (*tmallservice.TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse, error) {
	var resp tmallservice.TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

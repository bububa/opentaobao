package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallFuwuHomedecorationSupplyruleList 查询供给规则接口
// tmall.fuwu.homedecoration.supplyrule.list
//
// 查询供给规则接口
func TmallFuwuHomedecorationSupplyruleList(clt *core.SDKClient, req *tmallservice.TmallFuwuHomedecorationSupplyruleListAPIRequest, session string) (*tmallservice.TmallFuwuHomedecorationSupplyruleListAPIResponse, error) {
	var resp tmallservice.TmallFuwuHomedecorationSupplyruleListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

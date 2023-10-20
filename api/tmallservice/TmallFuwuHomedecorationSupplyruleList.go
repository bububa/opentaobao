package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallFuwuHomedecorationSupplyruleList 查询供给规则接口
// tmall.fuwu.homedecoration.supplyrule.list
//
// 查询供给规则接口
func TmallFuwuHomedecorationSupplyruleList(clt *core.SDKClient, req *tmallservice.TmallFuwuHomedecorationSupplyruleListAPIRequest, resp *tmallservice.TmallFuwuHomedecorationSupplyruleListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

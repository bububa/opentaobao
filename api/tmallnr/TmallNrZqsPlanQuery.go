package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrzqsplanquery 周期送配送明细查询
// tmall.nr.zqs.plan.query
//
// 周期送配送明细查询
func Tmallnrzqsplanquery(clt *core.SDKClient, req *tmallnr.TmallnrzqsplanqueryAPIRequest, session string) (*tmallnr.TmallnrzqsplanqueryAPIResponse, error) {
	var resp tmallnr.TmallnrzqsplanqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

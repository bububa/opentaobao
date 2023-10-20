package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrfulfillcancelreasonquery 查询取消履约的原因列表
// tmall.nr.fulfill.cancel.reason.query
//
// 新零售门店业务查询取消上门揽件的原因列表
func Tmallnrfulfillcancelreasonquery(clt *core.SDKClient, req *tmallnr.TmallnrfulfillcancelreasonqueryAPIRequest, session string) (*tmallnr.TmallnrfulfillcancelreasonqueryAPIResponse, error) {
	var resp tmallnr.TmallnrfulfillcancelreasonqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

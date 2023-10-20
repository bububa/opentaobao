package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrFulfillCancelReasonQuery 查询取消履约的原因列表
// tmall.nr.fulfill.cancel.reason.query
//
// 新零售门店业务查询取消上门揽件的原因列表
func TmallNrFulfillCancelReasonQuery(clt *core.SDKClient, req *tmallnr.TmallNrFulfillCancelReasonQueryAPIRequest, resp *tmallnr.TmallNrFulfillCancelReasonQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

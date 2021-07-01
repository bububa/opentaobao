package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

/* TmallNrFulfillCancel
取消上门揽件
tmall.nr.fulfill.cancel

新零售门店业务取消上门揽件 */
func TmallNrFulfillCancel(clt *core.SDKClient, req *tmallnr.TmallNrFulfillCancelAPIRequest, session string) (*tmallnr.TmallNrFulfillCancelAPIResponse, error) {
	var resp tmallnr.TmallNrFulfillCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

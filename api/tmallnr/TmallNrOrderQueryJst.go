package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

/* TmallNrOrderQueryJst
获取同城配送业务单笔订单
tmall.nr.order.query.jst

同城配送业务获取单笔订单 */
func TmallNrOrderQueryJst(clt *core.SDKClient, req *tmallnr.TmallNrOrderQueryJstAPIRequest, session string) (*tmallnr.TmallNrOrderQueryJstAPIResponse, error) {
	var resp tmallnr.TmallNrOrderQueryJstAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

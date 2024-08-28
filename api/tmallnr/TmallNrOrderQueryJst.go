package tmallnr

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrOrderQueryJst 获取同城配送业务单笔订单
// tmall.nr.order.query.jst
//
// 同城配送业务获取单笔订单
func TmallNrOrderQueryJst(ctx context.Context, clt *core.SDKClient, req *tmallnr.TmallNrOrderQueryJstAPIRequest, resp *tmallnr.TmallNrOrderQueryJstAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

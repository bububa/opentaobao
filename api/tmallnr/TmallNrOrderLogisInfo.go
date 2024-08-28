package tmallnr

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrOrderLogisInfo 区域零售订单获取取件码
// tmall.nr.order.logis.info
//
// 区域零售订单获取取件码，方便商家系统接入，获取取件码打印信息进行打印。
func TmallNrOrderLogisInfo(ctx context.Context, clt *core.SDKClient, req *tmallnr.TmallNrOrderLogisInfoAPIRequest, resp *tmallnr.TmallNrOrderLogisInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

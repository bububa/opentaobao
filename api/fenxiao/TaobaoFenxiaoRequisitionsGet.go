package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoRequisitionsGet 合作申请查询
// taobao.fenxiao.requisitions.get
//
// 合作申请查询
func TaobaoFenxiaoRequisitionsGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoRequisitionsGetAPIRequest, resp *fenxiao.TaobaoFenxiaoRequisitionsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

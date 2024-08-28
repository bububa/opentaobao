package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoOrderRemarkUpdate 修改采购单备注
// taobao.fenxiao.order.remark.update
//
// 供应商修改采购单备注
func TaobaoFenxiaoOrderRemarkUpdate(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoOrderRemarkUpdateAPIRequest, resp *fenxiao.TaobaoFenxiaoOrderRemarkUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

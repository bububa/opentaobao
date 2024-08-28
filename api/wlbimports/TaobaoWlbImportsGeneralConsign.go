package wlbimports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// TaobaoWlbImportsGeneralConsign 一般进口发货
// taobao.wlb.imports.general.consign
//
// 将订单信息发送到菜鸟海外转运仓；
// 业务规则：
// 1）交易订单为待发货状态。
// 2）单笔订单多个商品，交易金额不能大于1000人民币。
func TaobaoWlbImportsGeneralConsign(ctx context.Context, clt *core.SDKClient, req *wlbimports.TaobaoWlbImportsGeneralConsignAPIRequest, resp *wlbimports.TaobaoWlbImportsGeneralConsignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

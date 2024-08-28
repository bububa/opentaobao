package wms

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsCainiaoBillQuery 查询单据列表
// taobao.wlb.wms.cainiao.bill.query
//
// 查询单据列表
func TaobaoWlbWmsCainiaoBillQuery(ctx context.Context, clt *core.SDKClient, req *wms.TaobaoWlbWmsCainiaoBillQueryAPIRequest, resp *wms.TaobaoWlbWmsCainiaoBillQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

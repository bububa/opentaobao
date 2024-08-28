package xhotelonlineorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelCommoninvoiceListVtwo 用户常用发票信息查询接口
// taobao.xhotel.commoninvoice.list.vtwo
//
// 获取用户常用发票信息接口
func TaobaoXhotelCommoninvoiceListVtwo(ctx context.Context, clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelCommoninvoiceListVtwoAPIRequest, resp *xhotelonlineorder.TaobaoXhotelCommoninvoiceListVtwoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

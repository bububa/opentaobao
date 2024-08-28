package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TmallCityretailTxdFulfillOrderVirtualnumber 淘鲜达虚拟号服务接口
// tmall.cityretail.txd.fulfill.order.virtualnumber
//
// 虚拟小号绑定接口，只有开通了虚拟号服务的淘鲜达商家能绑定。
func TmallCityretailTxdFulfillOrderVirtualnumber(ctx context.Context, clt *core.SDKClient, req *wdk.TmallCityretailTxdFulfillOrderVirtualnumberAPIRequest, resp *wdk.TmallCityretailTxdFulfillOrderVirtualnumberAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

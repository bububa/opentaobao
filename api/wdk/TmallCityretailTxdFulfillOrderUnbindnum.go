package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TmallCityretailTxdFulfillOrderUnbindnum 淘鲜达虚拟号服务解绑接口
// tmall.cityretail.txd.fulfill.order.unbindnum
//
// 淘鲜达虚拟号解绑服务接口，通过订阅关系id进行解绑。
func TmallCityretailTxdFulfillOrderUnbindnum(ctx context.Context, clt *core.SDKClient, req *wdk.TmallCityretailTxdFulfillOrderUnbindnumAPIRequest, resp *wdk.TmallCityretailTxdFulfillOrderUnbindnumAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkOrderSyncWithitem 订单和商品同步接口
// alibaba.wdk.order.sync.withitem
//
// 轻pos,将订单和商品的信息一起传到盒马这边，进行创单和添加商品处理。
func AlibabaWdkOrderSyncWithitem(clt *core.SDKClient, req *wdk.AlibabaWdkOrderSyncWithitemAPIRequest, session string) (*wdk.AlibabaWdkOrderSyncWithitemAPIResponse, error) {
	var resp wdk.AlibabaWdkOrderSyncWithitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

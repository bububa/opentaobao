package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkOrderSync 五道口外部订单同步
// alibaba.wdk.order.sync
//
// 外部商户使用自助POS下单订单同步到五道口
func AlibabaWdkOrderSync(clt *core.SDKClient, req *wdk.AlibabaWdkOrderSyncAPIRequest, session string) (*wdk.AlibabaWdkOrderSyncAPIResponse, error) {
	var resp wdk.AlibabaWdkOrderSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

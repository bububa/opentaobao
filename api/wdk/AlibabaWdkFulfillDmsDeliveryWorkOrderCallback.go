package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkFulfillDmsDeliveryWorkOrderCallback 末端配配送作业回传
// alibaba.wdk.fulfill.dms.delivery.work.order.callback
//
// 末端配配送作业回传。
func AlibabaWdkFulfillDmsDeliveryWorkOrderCallback(clt *core.SDKClient, req *wdk.AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest, session string) (*wdk.AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse, error) {
	var resp wdk.AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

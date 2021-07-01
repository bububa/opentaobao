package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkFulfillWarehouseWorkOrderCallback
标准化仓作业单回传接口
alibaba.wdk.fulfill.warehouse.work.order.callback

标准化仓作业单回传接口 */
func AlibabaWdkFulfillWarehouseWorkOrderCallback(clt *core.SDKClient, req *wdk.AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest, session string) (*wdk.AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse, error) {
	var resp wdk.AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

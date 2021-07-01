package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallback
大润发B2C仓作业单回传接口
alibaba.wdk.fulfill.rt.btoc.warehouse.work.order.callback

大润发B2C仓作业单回传接口 */
func AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallback(clt *core.SDKClient, req *wdk.AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest, session string) (*wdk.AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse, error) {
	var resp wdk.AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Wdkwarehouseordercancel 仓作业取消下发
// wdk.warehouse.order.cancel
//
// 仓作业取消下发
func Wdkwarehouseordercancel(clt *core.SDKClient, req *wdk.WdkwarehouseordercancelAPIRequest, session string) (*wdk.WdkwarehouseordercancelAPIResponse, error) {
	var resp wdk.WdkwarehouseordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

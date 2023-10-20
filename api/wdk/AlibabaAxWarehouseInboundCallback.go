package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaaxwarehouseinboundcallback 翱象入库回传
// alibaba.ax.warehouse.inbound.callback
//
// 翱象入库回传
func Alibabaaxwarehouseinboundcallback(clt *core.SDKClient, req *wdk.AlibabaaxwarehouseinboundcallbackAPIRequest, session string) (*wdk.AlibabaaxwarehouseinboundcallbackAPIResponse, error) {
	var resp wdk.AlibabaaxwarehouseinboundcallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

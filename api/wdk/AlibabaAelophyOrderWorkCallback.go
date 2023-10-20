package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAelophyOrderWorkCallback 仓配作业结果回传接口
// alibaba.aelophy.order.work.callback
//
// 仓配作业结果回传接口
func AlibabaAelophyOrderWorkCallback(clt *core.SDKClient, req *wdk.AlibabaAelophyOrderWorkCallbackAPIRequest, session string) (*wdk.AlibabaAelophyOrderWorkCallbackAPIResponse, error) {
	var resp wdk.AlibabaAelophyOrderWorkCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaaelophyorderworkcallback 仓配作业结果回传接口
// alibaba.aelophy.order.work.callback
//
// 仓配作业结果回传接口
func Alibabaaelophyorderworkcallback(clt *core.SDKClient, req *wdk.AlibabaaelophyorderworkcallbackAPIRequest, session string) (*wdk.AlibabaaelophyorderworkcallbackAPIResponse, error) {
	var resp wdk.AlibabaaelophyorderworkcallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

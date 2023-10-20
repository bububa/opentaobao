package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkfinanceorderbackflow 财务订单回流
// alibaba.wdk.finance.order.backflow
//
// 星巴克拉取财务订单回流数据
func Alibabawdkfinanceorderbackflow(clt *core.SDKClient, req *wdk.AlibabawdkfinanceorderbackflowAPIRequest, session string) (*wdk.AlibabawdkfinanceorderbackflowAPIResponse, error) {
	var resp wdk.AlibabawdkfinanceorderbackflowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

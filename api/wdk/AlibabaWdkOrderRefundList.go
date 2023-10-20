package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkorderrefundlist 五道口交易退款批量查询
// alibaba.wdk.order.refund.list
//
// 按照条件查询退款数据
func Alibabawdkorderrefundlist(clt *core.SDKClient, req *wdk.AlibabawdkorderrefundlistAPIRequest, session string) (*wdk.AlibabawdkorderrefundlistAPIResponse, error) {
	var resp wdk.AlibabawdkorderrefundlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

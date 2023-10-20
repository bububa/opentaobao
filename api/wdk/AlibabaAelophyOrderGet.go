package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaaelophyorderget 翱象拉取订单接口
// alibaba.aelophy.order.get
//
// 翱象拉取订单接口
func Alibabaaelophyorderget(clt *core.SDKClient, req *wdk.AlibabaaelophyordergetAPIRequest, session string) (*wdk.AlibabaaelophyordergetAPIResponse, error) {
	var resp wdk.AlibabaaelophyordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

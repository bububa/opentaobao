package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaaelophyorderdelivererchange 配送员信息变更接口
// alibaba.aelophy.order.deliverer.change
//
// 配送员信息变更接口
func Alibabaaelophyorderdelivererchange(clt *core.SDKClient, req *wdk.AlibabaaelophyorderdelivererchangeAPIRequest, session string) (*wdk.AlibabaaelophyorderdelivererchangeAPIResponse, error) {
	var resp wdk.AlibabaaelophyorderdelivererchangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

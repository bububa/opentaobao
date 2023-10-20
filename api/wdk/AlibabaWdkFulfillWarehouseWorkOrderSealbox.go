package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkfulfillwarehouseworkordersealbox 仓封箱回告
// alibaba.wdk.fulfill.warehouse.work.order.sealbox
//
// 仓封箱回告箱与包裹的关系
func Alibabawdkfulfillwarehouseworkordersealbox(clt *core.SDKClient, req *wdk.AlibabawdkfulfillwarehouseworkordersealboxAPIRequest, session string) (*wdk.AlibabawdkfulfillwarehouseworkordersealboxAPIResponse, error) {
	var resp wdk.AlibabawdkfulfillwarehouseworkordersealboxAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

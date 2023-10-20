package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkoldposordercreate 淘鲜达外部商户老pos机产生的订单同步进淘鲜达
// alibaba.wdk.oldpos.order.create
//
// 淘鲜达外部商户老pos机产生的订单同步进淘鲜达
func Alibabawdkoldposordercreate(clt *core.SDKClient, req *wdk.AlibabawdkoldposordercreateAPIRequest, session string) (*wdk.AlibabawdkoldposordercreateAPIResponse, error) {
	var resp wdk.AlibabawdkoldposordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

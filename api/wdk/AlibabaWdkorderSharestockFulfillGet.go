package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkordersharestockfulfillget 商户订单履约数据获取
// alibaba.wdkorder.sharestock.fulfill.get
//
// 商户订单履约数据获取
func Alibabawdkordersharestockfulfillget(clt *core.SDKClient, req *wdk.AlibabawdkordersharestockfulfillgetAPIRequest, session string) (*wdk.AlibabawdkordersharestockfulfillgetAPIResponse, error) {
	var resp wdk.AlibabawdkordersharestockfulfillgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

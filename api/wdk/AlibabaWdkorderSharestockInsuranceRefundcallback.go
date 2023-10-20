package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkordersharestockinsurancerefundcallback 共享库存逆向订单理赔单回传
// alibaba.wdkorder.sharestock.insurance.refundcallback
//
// 共享库存逆向订单理赔单回传
func Alibabawdkordersharestockinsurancerefundcallback(clt *core.SDKClient, req *wdk.AlibabawdkordersharestockinsurancerefundcallbackAPIRequest, session string) (*wdk.AlibabawdkordersharestockinsurancerefundcallbackAPIResponse, error) {
	var resp wdk.AlibabawdkordersharestockinsurancerefundcallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

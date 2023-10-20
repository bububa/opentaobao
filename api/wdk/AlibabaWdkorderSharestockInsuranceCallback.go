package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkordersharestockinsurancecallback 共享库存订单投保后回传保单号
// alibaba.wdkorder.sharestock.insurance.callback
//
// 共享库存订单投保消息获取
func Alibabawdkordersharestockinsurancecallback(clt *core.SDKClient, req *wdk.AlibabawdkordersharestockinsurancecallbackAPIRequest, session string) (*wdk.AlibabawdkordersharestockinsurancecallbackAPIResponse, error) {
	var resp wdk.AlibabawdkordersharestockinsurancecallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

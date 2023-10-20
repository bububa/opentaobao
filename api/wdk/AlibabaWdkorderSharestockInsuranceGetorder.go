package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkordersharestockinsurancegetorder 共享库存订单投保消息获取
// alibaba.wdkorder.sharestock.insurance.getorder
//
// 共享库存订单投保消息获取
func Alibabawdkordersharestockinsurancegetorder(clt *core.SDKClient, req *wdk.AlibabawdkordersharestockinsurancegetorderAPIRequest, session string) (*wdk.AlibabawdkordersharestockinsurancegetorderAPIResponse, error) {
	var resp wdk.AlibabawdkordersharestockinsurancegetorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

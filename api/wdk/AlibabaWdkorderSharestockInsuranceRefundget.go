package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkordersharestockinsurancerefundget 共享库存投保业务售后逆向订单数据获取
// alibaba.wdkorder.sharestock.insurance.refundget
//
// 共享库存投保业务售后逆向订单数据获取
func Alibabawdkordersharestockinsurancerefundget(clt *core.SDKClient, req *wdk.AlibabawdkordersharestockinsurancerefundgetAPIRequest, session string) (*wdk.AlibabawdkordersharestockinsurancerefundgetAPIResponse, error) {
	var resp wdk.AlibabawdkordersharestockinsurancerefundgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

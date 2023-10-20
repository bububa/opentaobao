package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkordersharestockorderget 猫超商户订单拉取
// alibaba.wdkorder.sharestock.order.get
//
// 商户拉取猫超订单数据
func Alibabawdkordersharestockorderget(clt *core.SDKClient, req *wdk.AlibabawdkordersharestockordergetAPIRequest, session string) (*wdk.AlibabawdkordersharestockordergetAPIResponse, error) {
	var resp wdk.AlibabawdkordersharestockordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

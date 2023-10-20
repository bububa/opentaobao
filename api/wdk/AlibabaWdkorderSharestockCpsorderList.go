package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkordersharestockcpsorderlist cps正向分销订单批量回流
// alibaba.wdkorder.sharestock.cpsorder.list
//
// cps正向分销订单批量回流
func Alibabawdkordersharestockcpsorderlist(clt *core.SDKClient, req *wdk.AlibabawdkordersharestockcpsorderlistAPIRequest, session string) (*wdk.AlibabawdkordersharestockcpsorderlistAPIResponse, error) {
	var resp wdk.AlibabawdkordersharestockcpsorderlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

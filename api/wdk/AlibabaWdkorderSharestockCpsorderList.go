package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkorderSharestockCpsorderList cps正向分销订单批量回流
// alibaba.wdkorder.sharestock.cpsorder.list
//
// cps正向分销订单批量回流
func AlibabaWdkorderSharestockCpsorderList(clt *core.SDKClient, req *wdk.AlibabaWdkorderSharestockCpsorderListAPIRequest, session string) (*wdk.AlibabaWdkorderSharestockCpsorderListAPIResponse, error) {
	var resp wdk.AlibabaWdkorderSharestockCpsorderListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

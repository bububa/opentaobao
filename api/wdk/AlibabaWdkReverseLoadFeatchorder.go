package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkReverseLoadFeatchorder 取货详情
// alibaba.wdk.reverse.load.featchorder
//
// 取货详情
func AlibabaWdkReverseLoadFeatchorder(clt *core.SDKClient, req *wdk.AlibabaWdkReverseLoadFeatchorderAPIRequest, resp *wdk.AlibabaWdkReverseLoadFeatchorderAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

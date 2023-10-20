package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkReverseReversedetail 退款详情
// alibaba.wdk.reverse.reversedetail
//
// 退款详情
func AlibabaWdkReverseReversedetail(clt *core.SDKClient, req *wdk.AlibabaWdkReverseReversedetailAPIRequest, resp *wdk.AlibabaWdkReverseReversedetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkReverseApplyrefund 逆向申请接口
// alibaba.wdk.reverse.applyrefund
//
// 逆向渲染
func AlibabaWdkReverseApplyrefund(clt *core.SDKClient, req *wdk.AlibabaWdkReverseApplyrefundAPIRequest, resp *wdk.AlibabaWdkReverseApplyrefundAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

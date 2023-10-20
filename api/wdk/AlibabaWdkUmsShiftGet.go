package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkUmsShiftGet 移库单获取
// alibaba.wdk.ums.shift.get
//
// 移库单获取
func AlibabaWdkUmsShiftGet(clt *core.SDKClient, req *wdk.AlibabaWdkUmsShiftGetAPIRequest, resp *wdk.AlibabaWdkUmsShiftGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

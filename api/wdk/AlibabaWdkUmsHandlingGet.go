package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkUmsHandlingGet 加工单-回流单（新接口）
// alibaba.wdk.ums.handling.get
//
// 加工单-回流单（新接口）
func AlibabaWdkUmsHandlingGet(clt *core.SDKClient, req *wdk.AlibabaWdkUmsHandlingGetAPIRequest, resp *wdk.AlibabaWdkUmsHandlingGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

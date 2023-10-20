package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkTimeGet 获得当前系统时间
// alibaba.wdk.time.get
//
// 获得当前系统时间
func AlibabaWdkTimeGet(clt *core.SDKClient, req *wdk.AlibabaWdkTimeGetAPIRequest, resp *wdk.AlibabaWdkTimeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

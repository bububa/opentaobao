package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSeriesEdit 系列品变更-更新系列
// alibaba.wdk.series.edit
//
// 系列品变更-更新系列
func AlibabaWdkSeriesEdit(clt *core.SDKClient, req *wdk.AlibabaWdkSeriesEditAPIRequest, resp *wdk.AlibabaWdkSeriesEditAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

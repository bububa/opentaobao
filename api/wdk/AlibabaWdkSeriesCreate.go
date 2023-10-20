package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSeriesCreate 系列品变更-新增系列
// alibaba.wdk.series.create
//
// 系列品变更-新增系列
func AlibabaWdkSeriesCreate(clt *core.SDKClient, req *wdk.AlibabaWdkSeriesCreateAPIRequest, resp *wdk.AlibabaWdkSeriesCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

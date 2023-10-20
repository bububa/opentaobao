package xiamitrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamitrace"
)

// XiamiContentResourceActionReport 曲库开放平台内容行为上报接口
// xiami.content.resource.action.report
//
// 合作方对接入的曲库开放内容上报行为日志
func XiamiContentResourceActionReport(clt *core.SDKClient, req *xiamitrace.XiamiContentResourceActionReportAPIRequest, resp *xiamitrace.XiamiContentResourceActionReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

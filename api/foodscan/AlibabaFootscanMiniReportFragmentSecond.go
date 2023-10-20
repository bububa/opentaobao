package foodscan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/foodscan"
)

// AlibabaFootscanMiniReportFragmentSecond 第二只脚生成报告接口
// alibaba.footscan.mini.report.fragment.second
//
// 第二只脚生成报告接口
func AlibabaFootscanMiniReportFragmentSecond(clt *core.SDKClient, req *foodscan.AlibabaFootscanMiniReportFragmentSecondAPIRequest, resp *foodscan.AlibabaFootscanMiniReportFragmentSecondAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

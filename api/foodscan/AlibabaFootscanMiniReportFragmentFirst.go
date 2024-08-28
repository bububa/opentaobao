package foodscan

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/foodscan"
)

// AlibabaFootscanMiniReportFragmentFirst 第一只脚生成报告接口
// alibaba.footscan.mini.report.fragment.first
//
// 第一只脚生成报告接口
func AlibabaFootscanMiniReportFragmentFirst(ctx context.Context, clt *core.SDKClient, req *foodscan.AlibabaFootscanMiniReportFragmentFirstAPIRequest, resp *foodscan.AlibabaFootscanMiniReportFragmentFirstAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package foodscan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/foodscan"
)

/* AlibabaFootscanMiniReportFragmentFirst
第一只脚生成报告接口
alibaba.footscan.mini.report.fragment.first

第一只脚生成报告接口 */
func AlibabaFootscanMiniReportFragmentFirst(clt *core.SDKClient, req *foodscan.AlibabaFootscanMiniReportFragmentFirstAPIRequest, session string) (*foodscan.AlibabaFootscanMiniReportFragmentFirstAPIResponse, error) {
	var resp foodscan.AlibabaFootscanMiniReportFragmentFirstAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

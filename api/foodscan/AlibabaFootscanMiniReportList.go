package foodscan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/foodscan"
)

// AlibabaFootscanMiniReportList 查询报告列表
// alibaba.footscan.mini.report.list
//
// 查询报告列表
func AlibabaFootscanMiniReportList(clt *core.SDKClient, req *foodscan.AlibabaFootscanMiniReportListAPIRequest, session string) (*foodscan.AlibabaFootscanMiniReportListAPIResponse, error) {
	var resp foodscan.AlibabaFootscanMiniReportListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

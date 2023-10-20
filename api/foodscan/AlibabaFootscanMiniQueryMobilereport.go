package foodscan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/foodscan"
)

// AlibabaFootscanMiniQueryMobilereport 根据scanId查询报告
// alibaba.footscan.mini.query.mobilereport
//
// 根据scanId查询报告
func AlibabaFootscanMiniQueryMobilereport(clt *core.SDKClient, req *foodscan.AlibabaFootscanMiniQueryMobilereportAPIRequest, resp *foodscan.AlibabaFootscanMiniQueryMobilereportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

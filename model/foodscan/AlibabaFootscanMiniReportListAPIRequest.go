package foodscan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFootscanMiniReportListAPIRequest
查询报告列表 API请求
alibaba.footscan.mini.report.list

查询报告列表 */
type AlibabaFootscanMiniReportListAPIRequest struct {
	model.Params
	// 平台分配的token
	_token string
	// 请求数据
	_reqData *TobFeetModelMobileReportRequest
}

// New

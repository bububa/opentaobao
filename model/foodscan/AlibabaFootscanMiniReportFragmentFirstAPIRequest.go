package foodscan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFootscanMiniReportFragmentFirstAPIRequest
第一只脚生成报告接口 API请求
alibaba.footscan.mini.report.fragment.first

第一只脚生成报告接口 */
type AlibabaFootscanMiniReportFragmentFirstAPIRequest struct {
	model.Params
	// 平台分配的token
	_token string
	// 请求数据
	_reqData *FilePackageRequest
}

// New

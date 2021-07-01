package foodscan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFootscanMiniReportFragmentSecondAPIRequest
第二只脚生成报告接口 API请求
alibaba.footscan.mini.report.fragment.second

第二只脚生成报告接口 */
type AlibabaFootscanMiniReportFragmentSecondAPIRequest struct {
	model.Params
	// 平台分配的token
	_token string
	// 请求数据
	_reqData *FilePackageBasicReq
}

// New

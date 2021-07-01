package foodscan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFootscanMiniQueryMobilereportAPIRequest
根据scanId查询报告 API请求
alibaba.footscan.mini.query.mobilereport

根据scanId查询报告 */
type AlibabaFootscanMiniQueryMobilereportAPIRequest struct {
	model.Params
	// 平台分配的token
	_token string
	// 扫描ID
	_scanId string
}

// New

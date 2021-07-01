package foodscan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFootscanMiniSavedAPIRequest
更新报告状态 API请求
alibaba.footscan.mini.saved

更新报告状态接口 */
type AlibabaFootscanMiniSavedAPIRequest struct {
	model.Params
	// 平台分配的token
	_token string
	// 请求数据
	_reqData string
}

// New

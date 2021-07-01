package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNazcaAuthIssueauthapplyCallbackAPIRequest
出证申请回调 API请求
alibaba.nazca.auth.issueauthapply.callback

出证申请回调 */
type AlibabaNazcaAuthIssueauthapplyCallbackAPIRequest struct {
	model.Params
	// 合同编号
	_contractNum string
	// 出证机构
	_issueOrg string
	// 客户在1688的唯一标识
	_platformUserId string
	// 出证报告下载地址
	_reportUrl string
	// 出证状态
	_status string
}

// New

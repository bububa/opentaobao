package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqAppRisksummaryGetAPIRequest
应用风险概要信息查询接口 API请求
alibaba.security.jaq.app.risksummary.get

用户通过alibaba.security.jaq.app.risk.scan接口提交应用进行风险扫描后，用此接口获取风险概要信息，本接口不返回风险详细信息 */
type AlibabaSecurityJaqAppRisksummaryGetAPIRequest struct {
	model.Params
	// 任务唯一标识
	_itemId string
}

// New

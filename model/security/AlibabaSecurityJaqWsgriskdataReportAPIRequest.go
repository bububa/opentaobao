package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqWsgriskdataReportAPIRequest
无线保镖SDK风控数据上报 API请求
alibaba.security.jaq.wsgriskdata.report

无线保镖sdk根据用户的需要，上报数据到聚安全云端 */
type AlibabaSecurityJaqWsgriskdataReportAPIRequest struct {
	model.Params
	// wua串
	_wua string
	// mtopappkey是mtop的appkey
	_extParam string
}

// New

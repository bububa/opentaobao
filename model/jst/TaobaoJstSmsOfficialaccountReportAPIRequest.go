package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSmsOfficialaccountReportAPIRequest
聚石塔公众号信息上报 API请求
taobao.jst.sms.officialaccount.report

聚石塔公众号信息上报 */
type TaobaoJstSmsOfficialaccountReportAPIRequest struct {
	model.Params
	// 公众号信息上报接口入参
	_officialAccountInfoReportRequest *OfficialAccountInfoReportRequest
}

// New

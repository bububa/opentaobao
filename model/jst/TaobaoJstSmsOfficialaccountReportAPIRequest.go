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

// NewTaobaoJstSmsOfficialaccountReportRequest 初始化TaobaoJstSmsOfficialaccountReportAPIRequest对象
func NewTaobaoJstSmsOfficialaccountReportRequest() *TaobaoJstSmsOfficialaccountReportAPIRequest {
	return &TaobaoJstSmsOfficialaccountReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsOfficialaccountReportAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.officialaccount.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsOfficialaccountReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OfficialAccountInfoReportRequest Setter
// 公众号信息上报接口入参
func (r *TaobaoJstSmsOfficialaccountReportAPIRequest) SetOfficialAccountInfoReportRequest(_officialAccountInfoReportRequest *OfficialAccountInfoReportRequest) error {
	r._officialAccountInfoReportRequest = _officialAccountInfoReportRequest
	r.Set("official_account_info_report_request", _officialAccountInfoReportRequest)
	return nil
}

// Get OfficialAccountInfoReportRequest Getter
func (r TaobaoJstSmsOfficialaccountReportAPIRequest) GetOfficialAccountInfoReportRequest() *OfficialAccountInfoReportRequest {
	return r._officialAccountInfoReportRequest
}

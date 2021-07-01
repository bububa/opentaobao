package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpEffectAccountListAPIRequest
账户-报表 API请求
alibaba.scbp.effect.account.list

账户-报表,支持最近7天，最近30天，以及180天内时间区间。 */
type AlibabaScbpEffectAccountListAPIRequest struct {
	model.Params
	// AccountQuery
	_p4pAccountReportQuery *AccountQuery
}

// NewAlibabaScbpEffectAccountListRequest 初始化AlibabaScbpEffectAccountListAPIRequest对象
func NewAlibabaScbpEffectAccountListRequest() *AlibabaScbpEffectAccountListAPIRequest {
	return &AlibabaScbpEffectAccountListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpEffectAccountListAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.effect.account.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpEffectAccountListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is P4pAccountReportQuery Setter
// AccountQuery
func (r *AlibabaScbpEffectAccountListAPIRequest) SetP4pAccountReportQuery(_p4pAccountReportQuery *AccountQuery) error {
	r._p4pAccountReportQuery = _p4pAccountReportQuery
	r.Set("p4p_account_report_query", _p4pAccountReportQuery)
	return nil
}

// Get P4pAccountReportQuery Getter
func (r AlibabaScbpEffectAccountListAPIRequest) GetP4pAccountReportQuery() *AccountQuery {
	return r._p4pAccountReportQuery
}

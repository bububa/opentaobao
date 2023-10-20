package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpeffectaccountlistAPIRequest 账户-报表 API请求
// alibaba.scbp.effect.account.list
//
// 账户-报表,支持最近7天，最近30天，以及180天内时间区间。
type AlibabascbpeffectaccountlistAPIRequest struct {
	model.Params
	// AccountQuery
	_p4pAccountReportQuery *AccountQuery
}

// NewAlibabascbpeffectaccountlistRequest 初始化AlibabascbpeffectaccountlistAPIRequest对象
func NewAlibabascbpeffectaccountlistRequest() *AlibabascbpeffectaccountlistAPIRequest {
	return &AlibabascbpeffectaccountlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpeffectaccountlistAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.effect.account.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpeffectaccountlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpeffectaccountlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetP4pAccountReportQuery is P4pAccountReportQuery Setter
// AccountQuery
func (r *AlibabascbpeffectaccountlistAPIRequest) SetP4pAccountReportQuery(_p4pAccountReportQuery *AccountQuery) error {
	r._p4pAccountReportQuery = _p4pAccountReportQuery
	r.Set("p4p_account_report_query", _p4pAccountReportQuery)
	return nil
}

// GetP4pAccountReportQuery P4pAccountReportQuery Getter
func (r AlibabascbpeffectaccountlistAPIRequest) GetP4pAccountReportQuery() *AccountQuery {
	return r._p4pAccountReportQuery
}

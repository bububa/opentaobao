package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadreportgetaccountreportAPIRequest 账户报告 API请求
// alibaba.scbp.ad.report.get.account.report
//
// 账户报告
type AlibabascbpadreportgetaccountreportAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 请求参数
	_accountReportOperation *AccountReportOperationDto
}

// NewAlibabascbpadreportgetaccountreportRequest 初始化AlibabascbpadreportgetaccountreportAPIRequest对象
func NewAlibabascbpadreportgetaccountreportRequest() *AlibabascbpadreportgetaccountreportAPIRequest {
	return &AlibabascbpadreportgetaccountreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadreportgetaccountreportAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.report.get.account.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadreportgetaccountreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadreportgetaccountreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadreportgetaccountreportAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadreportgetaccountreportAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetAccountReportOperation is AccountReportOperation Setter
// 请求参数
func (r *AlibabascbpadreportgetaccountreportAPIRequest) SetAccountReportOperation(_accountReportOperation *AccountReportOperationDto) error {
	r._accountReportOperation = _accountReportOperation
	r.Set("account_report_operation", _accountReportOperation)
	return nil
}

// GetAccountReportOperation AccountReportOperation Getter
func (r AlibabascbpadreportgetaccountreportAPIRequest) GetAccountReportOperation() *AccountReportOperationDto {
	return r._accountReportOperation
}

package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdReportGetAccountReportAPIRequest 账户报告 API请求
// alibaba.scbp.ad.report.get.account.report
//
// 账户报告
type AlibabaScbpAdReportGetAccountReportAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 请求参数
	_accountReportOperation *AccountReportOperationDto
}

// NewAlibabaScbpAdReportGetAccountReportRequest 初始化AlibabaScbpAdReportGetAccountReportAPIRequest对象
func NewAlibabaScbpAdReportGetAccountReportRequest() *AlibabaScbpAdReportGetAccountReportAPIRequest {
	return &AlibabaScbpAdReportGetAccountReportAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdReportGetAccountReportAPIRequest) Reset() {
	r._topContext = nil
	r._accountReportOperation = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportGetAccountReportAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.report.get.account.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportGetAccountReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdReportGetAccountReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportGetAccountReportAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdReportGetAccountReportAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetAccountReportOperation is AccountReportOperation Setter
// 请求参数
func (r *AlibabaScbpAdReportGetAccountReportAPIRequest) SetAccountReportOperation(_accountReportOperation *AccountReportOperationDto) error {
	r._accountReportOperation = _accountReportOperation
	r.Set("account_report_operation", _accountReportOperation)
	return nil
}

// GetAccountReportOperation AccountReportOperation Getter
func (r AlibabaScbpAdReportGetAccountReportAPIRequest) GetAccountReportOperation() *AccountReportOperationDto {
	return r._accountReportOperation
}

var poolAlibabaScbpAdReportGetAccountReportAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdReportGetAccountReportRequest()
	},
}

// GetAlibabaScbpAdReportGetAccountReportRequest 从 sync.Pool 获取 AlibabaScbpAdReportGetAccountReportAPIRequest
func GetAlibabaScbpAdReportGetAccountReportAPIRequest() *AlibabaScbpAdReportGetAccountReportAPIRequest {
	return poolAlibabaScbpAdReportGetAccountReportAPIRequest.Get().(*AlibabaScbpAdReportGetAccountReportAPIRequest)
}

// ReleaseAlibabaScbpAdReportGetAccountReportAPIRequest 将 AlibabaScbpAdReportGetAccountReportAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdReportGetAccountReportAPIRequest(v *AlibabaScbpAdReportGetAccountReportAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdReportGetAccountReportAPIRequest.Put(v)
}

package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmMarketingIssueVoucherAPIRequest 发券 API请求
// alibaba.alsc.crm.marketing.issue.voucher
//
// 提供发券功能
type AlibabaAlscCrmMarketingIssueVoucherAPIRequest struct {
	model.Params
	// 参数
	_paramIssueVoucherReq *IssueVoucherReq
}

// NewAlibabaAlscCrmMarketingIssueVoucherRequest 初始化AlibabaAlscCrmMarketingIssueVoucherAPIRequest对象
func NewAlibabaAlscCrmMarketingIssueVoucherRequest() *AlibabaAlscCrmMarketingIssueVoucherAPIRequest {
	return &AlibabaAlscCrmMarketingIssueVoucherAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmMarketingIssueVoucherAPIRequest) Reset() {
	r._paramIssueVoucherReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmMarketingIssueVoucherAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.marketing.issue.voucher"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmMarketingIssueVoucherAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmMarketingIssueVoucherAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamIssueVoucherReq is ParamIssueVoucherReq Setter
// 参数
func (r *AlibabaAlscCrmMarketingIssueVoucherAPIRequest) SetParamIssueVoucherReq(_paramIssueVoucherReq *IssueVoucherReq) error {
	r._paramIssueVoucherReq = _paramIssueVoucherReq
	r.Set("param_issue_voucher_req", _paramIssueVoucherReq)
	return nil
}

// GetParamIssueVoucherReq ParamIssueVoucherReq Getter
func (r AlibabaAlscCrmMarketingIssueVoucherAPIRequest) GetParamIssueVoucherReq() *IssueVoucherReq {
	return r._paramIssueVoucherReq
}

var poolAlibabaAlscCrmMarketingIssueVoucherAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmMarketingIssueVoucherRequest()
	},
}

// GetAlibabaAlscCrmMarketingIssueVoucherRequest 从 sync.Pool 获取 AlibabaAlscCrmMarketingIssueVoucherAPIRequest
func GetAlibabaAlscCrmMarketingIssueVoucherAPIRequest() *AlibabaAlscCrmMarketingIssueVoucherAPIRequest {
	return poolAlibabaAlscCrmMarketingIssueVoucherAPIRequest.Get().(*AlibabaAlscCrmMarketingIssueVoucherAPIRequest)
}

// ReleaseAlibabaAlscCrmMarketingIssueVoucherAPIRequest 将 AlibabaAlscCrmMarketingIssueVoucherAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmMarketingIssueVoucherAPIRequest(v *AlibabaAlscCrmMarketingIssueVoucherAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmMarketingIssueVoucherAPIRequest.Put(v)
}

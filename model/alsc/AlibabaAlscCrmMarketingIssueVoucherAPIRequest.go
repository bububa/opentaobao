package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmmarketingissuevoucherAPIRequest 发券 API请求
// alibaba.alsc.crm.marketing.issue.voucher
//
// 提供发券功能
type AlibabaalsccrmmarketingissuevoucherAPIRequest struct {
	model.Params
	// 参数
	_paramIssueVoucherReq *IssueVoucherReq
}

// NewAlibabaalsccrmmarketingissuevoucherRequest 初始化AlibabaalsccrmmarketingissuevoucherAPIRequest对象
func NewAlibabaalsccrmmarketingissuevoucherRequest() *AlibabaalsccrmmarketingissuevoucherAPIRequest {
	return &AlibabaalsccrmmarketingissuevoucherAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmmarketingissuevoucherAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.marketing.issue.voucher"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmmarketingissuevoucherAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmmarketingissuevoucherAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamIssueVoucherReq is ParamIssueVoucherReq Setter
// 参数
func (r *AlibabaalsccrmmarketingissuevoucherAPIRequest) SetParamIssueVoucherReq(_paramIssueVoucherReq *IssueVoucherReq) error {
	r._paramIssueVoucherReq = _paramIssueVoucherReq
	r.Set("param_issue_voucher_req", _paramIssueVoucherReq)
	return nil
}

// GetParamIssueVoucherReq ParamIssueVoucherReq Getter
func (r AlibabaalsccrmmarketingissuevoucherAPIRequest) GetParamIssueVoucherReq() *IssueVoucherReq {
	return r._paramIssueVoucherReq
}

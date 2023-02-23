package alitripreceipt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripReceiptIssueresultNotifyAPIRequest 飞猪发票开票结果通知 API请求
// alitrip.receipt.issueresult.notify
//
// 飞猪发票开票结果通知
type AlitripReceiptIssueresultNotifyAPIRequest struct {
	model.Params
	// 开票结果通知
	_param0 *IssueResultNotifyCmd
}

// NewAlitripReceiptIssueresultNotifyRequest 初始化AlitripReceiptIssueresultNotifyAPIRequest对象
func NewAlitripReceiptIssueresultNotifyRequest() *AlitripReceiptIssueresultNotifyAPIRequest {
	return &AlitripReceiptIssueresultNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripReceiptIssueresultNotifyAPIRequest) GetApiMethodName() string {
	return "alitrip.receipt.issueresult.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripReceiptIssueresultNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripReceiptIssueresultNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 开票结果通知
func (r *AlitripReceiptIssueresultNotifyAPIRequest) SetParam0(_param0 *IssueResultNotifyCmd) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlitripReceiptIssueresultNotifyAPIRequest) GetParam0() *IssueResultNotifyCmd {
	return r._param0
}

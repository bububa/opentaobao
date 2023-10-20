package alitripreceipt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripreceiptissueresultnotifyAPIRequest 飞猪发票开票结果通知 API请求
// alitrip.receipt.issueresult.notify
//
// 飞猪发票开票结果通知
type AlitripreceiptissueresultnotifyAPIRequest struct {
	model.Params
	// 开票结果通知
	_param0 *IssueResultNotifyCmd
}

// NewAlitripreceiptissueresultnotifyRequest 初始化AlitripreceiptissueresultnotifyAPIRequest对象
func NewAlitripreceiptissueresultnotifyRequest() *AlitripreceiptissueresultnotifyAPIRequest {
	return &AlitripreceiptissueresultnotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripreceiptissueresultnotifyAPIRequest) GetApiMethodName() string {
	return "alitrip.receipt.issueresult.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripreceiptissueresultnotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripreceiptissueresultnotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 开票结果通知
func (r *AlitripreceiptissueresultnotifyAPIRequest) SetParam0(_param0 *IssueResultNotifyCmd) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlitripreceiptissueresultnotifyAPIRequest) GetParam0() *IssueResultNotifyCmd {
	return r._param0
}

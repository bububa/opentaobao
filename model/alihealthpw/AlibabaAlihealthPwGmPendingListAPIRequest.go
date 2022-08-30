package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwGmPendingListAPIRequest 同情用药待审核工单查询接口 API请求
// alibaba.alihealth.pw.gm.pending.list
//
// 同情用药待审核工单查询接口，提供给合作方用来查询待处理工单列表
type AlibabaAlihealthPwGmPendingListAPIRequest struct {
	model.Params
	// 入参
	_body *PendingListReq
}

// NewAlibabaAlihealthPwGmPendingListRequest 初始化AlibabaAlihealthPwGmPendingListAPIRequest对象
func NewAlibabaAlihealthPwGmPendingListRequest() *AlibabaAlihealthPwGmPendingListAPIRequest {
	return &AlibabaAlihealthPwGmPendingListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwGmPendingListAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.gm.pending.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwGmPendingListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBody is Body Setter
// 入参
func (r *AlibabaAlihealthPwGmPendingListAPIRequest) SetBody(_body *PendingListReq) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaAlihealthPwGmPendingListAPIRequest) GetBody() *PendingListReq {
	return r._body
}

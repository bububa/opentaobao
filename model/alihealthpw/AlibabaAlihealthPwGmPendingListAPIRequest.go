package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpwgmpendinglistAPIRequest 同情用药待审核工单查询接口 API请求
// alibaba.alihealth.pw.gm.pending.list
//
// 同情用药待审核工单查询接口，提供给合作方用来查询待处理工单列表
type AlibabaalihealthpwgmpendinglistAPIRequest struct {
	model.Params
	// 入参
	_body *PendingListReq
}

// NewAlibabaalihealthpwgmpendinglistRequest 初始化AlibabaalihealthpwgmpendinglistAPIRequest对象
func NewAlibabaalihealthpwgmpendinglistRequest() *AlibabaalihealthpwgmpendinglistAPIRequest {
	return &AlibabaalihealthpwgmpendinglistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthpwgmpendinglistAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.gm.pending.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthpwgmpendinglistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthpwgmpendinglistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 入参
func (r *AlibabaalihealthpwgmpendinglistAPIRequest) SetBody(_body *PendingListReq) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaalihealthpwgmpendinglistAPIRequest) GetBody() *PendingListReq {
	return r._body
}

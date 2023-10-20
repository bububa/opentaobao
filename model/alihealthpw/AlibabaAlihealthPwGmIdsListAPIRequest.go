package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpwgmidslistAPIRequest 同情用药根据申请单列表查询申请单 API请求
// alibaba.alihealth.pw.gm.ids.list
//
// 同情用药根据申请单列表查询申请单
type AlibabaalihealthpwgmidslistAPIRequest struct {
	model.Params
	// 入参
	_body *ListByApplyIdsForBreq
}

// NewAlibabaalihealthpwgmidslistRequest 初始化AlibabaalihealthpwgmidslistAPIRequest对象
func NewAlibabaalihealthpwgmidslistRequest() *AlibabaalihealthpwgmidslistAPIRequest {
	return &AlibabaalihealthpwgmidslistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthpwgmidslistAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.gm.ids.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthpwgmidslistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthpwgmidslistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 入参
func (r *AlibabaalihealthpwgmidslistAPIRequest) SetBody(_body *ListByApplyIdsForBreq) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaalihealthpwgmidslistAPIRequest) GetBody() *ListByApplyIdsForBreq {
	return r._body
}

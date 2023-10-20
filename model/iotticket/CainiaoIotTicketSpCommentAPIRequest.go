package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoiotticketspcommentAPIRequest IoT售后服务商工单备注 API请求
// cainiao.iot.ticket.sp.comment
//
// IoT售后服务商工单备注
type CainiaoiotticketspcommentAPIRequest struct {
	model.Params
	// 请求参数
	_param *CommentTicketTopRequest
}

// NewCainiaoiotticketspcommentRequest 初始化CainiaoiotticketspcommentAPIRequest对象
func NewCainiaoiotticketspcommentRequest() *CainiaoiotticketspcommentAPIRequest {
	return &CainiaoiotticketspcommentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoiotticketspcommentAPIRequest) GetApiMethodName() string {
	return "cainiao.iot.ticket.sp.comment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoiotticketspcommentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoiotticketspcommentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求参数
func (r *CainiaoiotticketspcommentAPIRequest) SetParam(_param *CommentTicketTopRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r CainiaoiotticketspcommentAPIRequest) GetParam() *CommentTicketTopRequest {
	return r._param
}
